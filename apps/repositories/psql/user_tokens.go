package psql

import (
	"context"
	"errors"
	"time"

	"multifinancetest/apps/models"

	"github.com/vizucode/gokit/logger"
	"gorm.io/gorm"
)

func (psql *psql) CreateUserToken(ctx context.Context, payload models.Tokens) (err error) {

	if err = psql.db.
		Model(&models.Tokens{}).
		WithContext(ctx).
		Create(&payload).Error; err != nil {
		logger.Log.Error(ctx, err)
		return err
	}

	return nil
}

func (psql *psql) UpdateUserTokenByRefreshToken(ctx context.Context, refreshToken string, payload models.Tokens) (err error) {

	selectedFile := []string{}

	if payload.AccessToken != "" {
		selectedFile = append(selectedFile, "access_token")
	}

	if payload.RefreshToken != "" {
		selectedFile = append(selectedFile, "refresh_token")
	}

	if payload.AccessTokenRevoked {
		selectedFile = append(selectedFile, "access_token_revoked")
	}

	if payload.RefreshTokenRevoked {
		selectedFile = append(selectedFile, "refresh_token_revoked")
	}

	if payload.AccessTokenExpiredAt != (time.Time{}) {
		selectedFile = append(selectedFile, "access_token_expired_at")
	}

	if payload.RefreshTokenExpiredAt != (time.Time{}) {
		selectedFile = append(selectedFile, "refresh_token_expired_at")
	}

	err = psql.db.
		Model(&models.Tokens{}).
		WithContext(ctx).
		Where("refresh_token = ?", refreshToken).
		Select(selectedFile).
		Updates(&payload).Error

	if err != nil {
		logger.Log.Error(ctx, err)
		return err
	}

	return nil
}

func (psql *psql) FirstActiveUserTokenByAccessToken(ctx context.Context, accessToken string) (resp models.Tokens, err error) {

	err = psql.db.
		Model(&models.Tokens{}).
		WithContext(ctx).
		Where("access_token = ?", accessToken).
		Where("access_token_revoked = ?", false).
		Where("refresh_token_revoked = ?", false).
		Where("access_token_expired_at > ?", time.Now().UTC()).
		Where("refresh_token_expired_at > ?", time.Now().UTC()).
		Order("updated_at desc").
		First(&resp).Error

	if err != nil {
		logger.Log.Error(ctx, err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return resp, nil
		}
		return resp, err
	}

	return resp, nil
}

func (psql *psql) FirstActiveUserTokenByUserId(ctx context.Context, userId string) (resp models.Tokens, err error) {
	err = psql.db.
		Model(&models.Tokens{}).
		WithContext(ctx).
		Where("customer_id = ?", userId).
		Where("access_token_revoked = ?", false).
		Where("refresh_token_revoked = ?", false).
		Where("access_token_expired_at > ?", time.Now().UTC()).
		Where("refresh_token_expired_at > ?", time.Now().UTC()).
		Order("created_at desc").
		First(&resp).Error

	if err != nil {
		logger.Log.Error(ctx, err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return resp, nil
		}
		return resp, err
	}

	return resp, nil
}

func (psql *psql) RevokeAllTokenByAccessTokenByUserId(ctx context.Context, userId string) (err error) {

	selectedField := []string{
		"access_token_revoked",
		"refresh_token_revoked",
		"access_token_expired_at",
		"refresh_token_expired_at",
		"updated_at",
	}

	userToken := models.Tokens{
		AccessTokenRevoked:    true,
		RefreshTokenRevoked:   true,
		AccessTokenExpiredAt:  time.Now().Add(-1 * time.Hour).UTC(),
		RefreshTokenExpiredAt: time.Now().Add(-1 * time.Hour).UTC(),
		UpdatedAt:             time.Now().UTC(),
	}

	err = psql.db.
		Model(&models.Tokens{}).
		WithContext(ctx).
		Where("customer_id = ?", userId).
		Select(selectedField).
		Updates(&userToken).Error

	if err != nil {
		logger.Log.Error(ctx, err)
		return err
	}

	return nil
}

func (psql *psql) FirstActiveRefreshToken(ctx context.Context, refreshToken string) (resp models.Tokens, err error) {
	err = psql.db.
		Model(&models.Tokens{}).
		WithContext(ctx).
		Where("refresh_token = ?", refreshToken).
		Where("refresh_token_revoked = ?", false).
		Where("refresh_token_expired_at > ?", time.Now().UTC()).
		First(&resp).Error

	return resp, nil
}
