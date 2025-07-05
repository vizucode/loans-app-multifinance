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
