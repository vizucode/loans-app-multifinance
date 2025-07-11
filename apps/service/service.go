package service

import (
	"context"

	"multifinancetest/apps/domain"
)

type IAuthService interface {
	SignUp(ctx context.Context, req domain.RequestCustomer) (err error)
	SignIn(ctx context.Context, req domain.RequestSignIn) (resp domain.ResponseSignIn, err error)
	SignOut(ctx context.Context) (err error)
	FirstCustomer(ctx context.Context) (resp domain.ResponseCustomer, err error)
	RefreshToken(ctx context.Context, accessToken string, refreshToken string) (resp domain.ResponseSignIn, err error)
}

type ILoans interface {
	CreateLoan(ctx context.Context, req domain.RequestLoans) (resp domain.ResponseLoans, err error)
	GetLimitLoans(ctx context.Context) (resp []domain.ResponseLimitLoans, err error)
	GetHistoryLoans(ctx context.Context, filter domain.Filter) (resp []domain.ResponseHistoryLoans, err error)
}
