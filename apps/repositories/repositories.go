package repositories

import (
	"context"

	"multifinancetest/apps/models"
)

type IDatabase interface {
	// Customers
	CreateCustomer(ctx context.Context, payload models.Customer) (err error)
	FirstCustomerById(ctx context.Context, id string) (resp models.Customer, err error)
	FirstCustomerByEmail(ctx context.Context, email string) (resp models.Customer, err error)

	// Token
	CreateUserToken(ctx context.Context, payload models.Tokens) (err error)
	UpdateUserTokenByRefreshToken(ctx context.Context, refreshToken string, payload models.Tokens) (err error)
	FirstActiveRefreshToken(ctx context.Context, refreshToken string) (resp models.Tokens, err error)
	FirstActiveUserTokenByAccessToken(ctx context.Context, accessToken string) (resp models.Tokens, err error)
	FirstActiveUserTokenByUserId(ctx context.Context, userId string) (resp models.Tokens, err error)
	RevokeAllTokenByAccessTokenByUserId(ctx context.Context, userId string) (err error)
}

type IStorage interface {
	UploadFile(ctx context.Context, filename string, file []byte) (err error)
	DeleteFile(ctx context.Context, key string) error
}
