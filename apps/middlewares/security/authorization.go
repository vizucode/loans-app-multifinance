package security

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"multifinancetest/apps/domain"
	contextkeys "multifinancetest/helpers/constants/context_keys"
	"multifinancetest/helpers/constants/httpstd"
	"multifinancetest/helpers/constants/rpcstd"

	"github.com/vizucode/gokit/logger"
	"github.com/vizucode/gokit/utils/env"
	"github.com/vizucode/gokit/utils/errorkit"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func (mw *security) AuthMiddleware(c *fiber.Ctx) error {

	var (
		UserContext domain.UserContext
	)

	ctx := c.Context()
	authHeader := c.Get("Authorization")

	if strings.EqualFold(authHeader, "") {
		return c.Next()
	}

	if !strings.HasPrefix(authHeader, "Bearer ") {
		return errorkit.NewErrorStd(http.StatusUnauthorized, rpcstd.INVALID_ARGUMENT, "Invalid token format")
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")

	if strings.EqualFold(token, "") {
		return c.Next()
	}

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signin method: %v", token)
		}
		return []byte(env.GetString("ACCESS_TOKEN_SECRET")), nil
	})

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		logger.Log.Error(ctx, err)
		return errorkit.NewErrorStd(http.StatusUnauthorized, rpcstd.INVALID_ARGUMENT, "token is expired")
	}

	exp, ok := claims["exp"].(float64)
	if !ok {
		logger.Log.Error(ctx, err)
		return errorkit.NewErrorStd(http.StatusUnauthorized, rpcstd.INVALID_ARGUMENT, "Invalid token claims")
	}

	expirationTime := time.Unix(int64(exp), 0).UTC()
	durationUntilExpiration := expirationTime.Sub(time.Now().UTC())

	if durationUntilExpiration <= 0 {
		err := errorkit.NewErrorStd(http.StatusUnauthorized, rpcstd.INVALID_ARGUMENT, "Token has already expired")
		logger.Log.Error(ctx, err)
		return errorkit.NewErrorStd(http.StatusUnauthorized, rpcstd.INVALID_ARGUMENT, "Token has already expired")
	}

	// Parsing user context from token
	resultClaims, err := json.Marshal(&claims)
	if err != nil {
		logger.Log.Error(ctx, err)
		return errorkit.NewErrorStd(http.StatusInternalServerError, rpcstd.INTERNAL, httpstd.InternalServerError)
	}

	err = json.Unmarshal(resultClaims, &UserContext)
	if err != nil {
		fmt.Println(err)
		return errorkit.NewErrorStd(http.StatusInternalServerError, rpcstd.INTERNAL, httpstd.InternalServerError)
	}

	// check to database if the token is valid
	resultUserToken, err := mw.db.FirstActiveUserTokenByAccessToken(ctx, token)
	if err != nil {
		logger.Log.Error(ctx, err)
		return errorkit.NewErrorStd(http.StatusUnauthorized, rpcstd.INVALID_ARGUMENT, "you have been sign out, please sign in again")
	}

	if strings.EqualFold(resultUserToken.AccessToken, "") {
		return errorkit.NewErrorStd(http.StatusUnauthorized, rpcstd.INVALID_ARGUMENT, "you have been sign out, please sign in again")
	}

	c.Locals(contextkeys.UserContext, UserContext)

	return c.Next()
}
