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

	// customers loan
	CreateCustomerTenor(ctx context.Context, payload models.CustomerTenors) (err error)
	GetAllCustomerTenor(ctx context.Context, customerID string) (resp []models.CustomerTenors, err error)
	GetCustomerLimitTenor(ctx context.Context, customerID string, tenorMonth int) (resp models.CustomerTenors, err error)

	// Tenors
	GetAllTenor(ctx context.Context) (resp []models.Tenor, err error)

	// Customer Loans
	CreateCustomerLoan(ctx context.Context, payload models.CustomerLoans) (err error)
	CreateTrxCustomerLoan(ctx context.Context, customerID string, tenorMonth int, payload models.CustomerLoans) (err error)
	CountLimitRemainingTenorMonthLoan(ctx context.Context, customerID string, tenorMonth int) (resp int, err error)
	GetCustomerLoans(ctx context.Context, customerId string, filter models.Filter) (resp []models.CustomerLoans, err error)

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
