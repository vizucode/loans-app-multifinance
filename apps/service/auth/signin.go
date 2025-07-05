package auth

import (
	"context"
	"net/http"
	"time"

	"multifinancetest/apps/domain"
	"multifinancetest/apps/models"
	"multifinancetest/helpers/constants/rpcstd"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/vizucode/gokit/logger"
	"github.com/vizucode/gokit/utils/env"
	"github.com/vizucode/gokit/utils/errorkit"
	"golang.org/x/crypto/bcrypt"
)

func (uc *auth) SignIn(ctx context.Context, req domain.RequestSignIn) (resp domain.ResponseSignIn, err error) {

	err = uc.validator.StructCtx(ctx, req)
	if err != nil {
		logger.Log.Error(ctx, err)
		return resp, err
	}

	resultUser, err := uc.db.FirstCustomerByEmail(ctx, req.Email)
	if err != nil {
		logger.Log.Error(ctx, err)
		return resp, errorkit.NewErrorStd(http.StatusBadRequest, rpcstd.ABORTED, "email not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(resultUser.Password), []byte(req.Password))
	if err != nil {
		logger.Log.Error(ctx, err)
		return resp, errorkit.NewErrorStd(http.StatusBadRequest, rpcstd.ABORTED, "incorrect password")
	}

	accessTokenExpired := time.Now().UTC().Add(time.Minute * 30)
	accessTokenClaims := jwt.MapClaims{
		"id":       string(resultUser.ID),
		"fullname": resultUser.FullName,
		"exp":      accessTokenExpired.Unix(),
	}

	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims).SignedString([]byte(env.GetString("ACCESS_TOKEN_SECRET")))
	if err != nil {
		logger.Log.Error(ctx, err)
		return resp, errorkit.NewErrorStd(http.StatusInternalServerError, rpcstd.ABORTED, "failed to generate token")
	}

	refreshTokenExpired := time.Now().UTC().Add(time.Hour * 24 * 30)
	refreshTokenClaims := jwt.MapClaims{
		"id":       string(resultUser.ID),
		"fullname": resultUser.FullName,
		"exp":      refreshTokenExpired.Unix(),
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims).SignedString([]byte(env.GetString("REFRESH_TOKEN_SECRET")))
	if err != nil {
		logger.Log.Error(ctx, err)
		return resp, errorkit.NewErrorStd(http.StatusInternalServerError, rpcstd.ABORTED, "failed to generate token")
	}

	// revoke all user token
	err = uc.db.RevokeAllTokenByAccessTokenByUserId(ctx, string(resultUser.ID))
	if err != nil {
		logger.Log.Error(ctx, err)
		return resp, err
	}

	defaultUUID := uuid.New()
	err = uc.db.CreateUserToken(ctx, models.Tokens{
		ID:                    defaultUUID.String(),
		CustomerId:            resultUser.ID,
		AccessToken:           accessToken,
		RefreshToken:          refreshToken,
		AccessTokenRevoked:    false,
		RefreshTokenRevoked:   false,
		AccessTokenExpiredAt:  accessTokenExpired,
		RefreshTokenExpiredAt: refreshTokenExpired,
	})
	if err != nil {
		logger.Log.Error(ctx, err)
		return resp, err
	}

	resp.AccessToken = accessToken
	resp.RefreshToken = refreshToken
	resp.AccessTokenExpired = accessTokenExpired.Format(time.RFC3339)

	return resp, nil
}
