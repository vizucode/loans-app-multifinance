package auth

import (
	"context"
	"net/http"
	"strings"
	"time"

	"multifinancetest/apps/domain"
	"multifinancetest/apps/models"
	"multifinancetest/helpers/constants/rpcstd"

	"github.com/golang-jwt/jwt/v4"
	"github.com/vizucode/gokit/logger"
	"github.com/vizucode/gokit/utils/env"
	"github.com/vizucode/gokit/utils/errorkit"
)

func (uc *auth) RefreshToken(ctx context.Context, accessToken string, refreshToken string) (resp domain.ResponseSignIn, err error) {
	resultUserToken, err := uc.db.FirstActiveRefreshToken(ctx, refreshToken)
	if err != nil {
		logger.Log.Error(ctx, err)
		return resp, errorkit.NewErrorStd(http.StatusBadRequest, rpcstd.ABORTED, "token not found")
	}

	if strings.EqualFold(resultUserToken.RefreshToken, "") {
		return resp, errorkit.NewErrorStd(http.StatusBadRequest, rpcstd.ABORTED, "token not found")
	}

	// get current user
	resultUser, err := uc.db.FirstCustomerById(ctx, string(resultUserToken.CustomerId))
	if err != nil {
		logger.Log.Error(ctx, err)
		return resp, errorkit.NewErrorStd(http.StatusBadRequest, rpcstd.ABORTED, "token not found")
	}

	_, err = uc.checkRefreshExpired(refreshToken)
	if err != nil {
		logger.Log.Error(ctx, err)
		return resp, err
	}

	// generate new access token
	accessTokenExpired := time.Now().UTC().Add(time.Minute * 30)
	accessTokenClaims := jwt.MapClaims{
		"id":       string(resultUser.ID),
		"fullname": resultUser.FullName,
		"exp":      accessTokenExpired.Unix(),
	}

	accessToken1, err := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims).SignedString([]byte(env.GetString("ACCESS_TOKEN_SECRET")))
	if err != nil {
		logger.Log.Error(ctx, err)
		return resp, errorkit.NewErrorStd(http.StatusInternalServerError, rpcstd.ABORTED, "failed to generate token")
	}

	// revoke all token
	err = uc.db.RevokeAllTokenByAccessTokenByUserId(ctx, string(resultUser.ID))
	if err != nil {
		logger.Log.Error(ctx, err)
		return resp, err
	}

	err = uc.db.CreateUserToken(ctx, models.Tokens{
		CustomerId:            resultUser.ID,
		AccessToken:           accessToken1,
		AccessTokenExpiredAt:  accessTokenExpired,
		RefreshToken:          refreshToken,
		RefreshTokenExpiredAt: resultUserToken.RefreshTokenExpiredAt,
	})
	if err != nil {
		logger.Log.Error(ctx, err)
		return resp, err
	}

	resp.AccessToken = accessToken1
	resp.RefreshToken = refreshToken
	resp.AccessTokenExpired = accessTokenExpired.Format(time.RFC3339)

	return resp, nil
}

func (uc *auth) checkRefreshExpired(refreshToken string) (isExpired bool, err error) {
	// check if refresh token wasn't expired
	refreshTokenClaims, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errorkit.NewErrorStd(http.StatusBadRequest, rpcstd.ABORTED, "signature refresh token not verified")
		}

		return []byte(env.GetString("REFRESH_TOKEN_SECRET")), nil
	})
	if err != nil {
		return isExpired, errorkit.NewErrorStd(http.StatusBadRequest, rpcstd.ABORTED, "signature refresh token not verified")
	}

	claims, ok := refreshTokenClaims.Claims.(jwt.MapClaims)
	if !ok {
		return isExpired, errorkit.NewErrorStd(http.StatusBadRequest, rpcstd.ABORTED, "signature refresh token not verified")
	}

	exp, ok := claims["exp"].(float64)
	if !ok {
		return isExpired, errorkit.NewErrorStd(http.StatusBadRequest, rpcstd.ABORTED, "signature refresh token not verified")
	}

	if int64(exp) < time.Now().UTC().Unix() {
		return isExpired, errorkit.NewErrorStd(http.StatusBadRequest, rpcstd.ABORTED, "refresh token was expired, please sign in again")
	}

	return false, nil
}
