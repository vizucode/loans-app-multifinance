package psqlmock

import (
	"context"
	"multifinancetest/apps/models"

	"github.com/stretchr/testify/mock"
)

type psqlMockRepo struct {
	mock.Mock
}

func NewPsqlMockRepo() *psqlMockRepo {
	return &psqlMockRepo{}
}

// ----------------------
// Customers
// ----------------------

func (m *psqlMockRepo) CreateCustomer(ctx context.Context, payload models.Customer) (err error) {
	return
}

func (m *psqlMockRepo) FirstCustomerById(ctx context.Context, id string) (resp models.Customer, err error) {
	return
}

func (m *psqlMockRepo) FirstCustomerByEmail(ctx context.Context, email string) (resp models.Customer, err error) {
	return
}

// ----------------------
// Customer Loans (Tenors)
// ----------------------

func (m *psqlMockRepo) CreateCustomerTenor(ctx context.Context, payload models.CustomerTenors) (err error) {
	return
}

func (m *psqlMockRepo) GetAllCustomerTenor(ctx context.Context, customerID string) (resp []models.CustomerTenors, err error) {
	return
}

func (m *psqlMockRepo) GetCustomerLimitTenor(ctx context.Context, customerID string, tenorMonth int) (resp models.CustomerTenors, err error) {
	return
}

// ----------------------
// Tenors
// ----------------------

func (m *psqlMockRepo) GetAllTenor(ctx context.Context) (resp []models.Tenor, err error) {
	return
}

// ----------------------
// Customer Loans
// ----------------------

func (m *psqlMockRepo) CreateCustomerLoan(ctx context.Context, payload models.CustomerLoans) (err error) {
	return
}

func (m *psqlMockRepo) CreateTrxCustomerLoan(ctx context.Context, customerID string, tenorMonth int, payload models.CustomerLoans) (err error) {

	args := m.Called(ctx, customerID, tenorMonth, payload)

	if response := args.Get(0); response != nil {
		return args.Error(0)
	}

	return nil
}

func (m *psqlMockRepo) CountLimitRemainingTenorMonthLoan(ctx context.Context, customerID string, tenorMonth int) (resp int, err error) {
	return
}

func (m *psqlMockRepo) GetCustomerLoans(ctx context.Context, customerId string, filter models.Filter) (resp []models.CustomerLoans, err error) {
	return
}

func (m *psqlMockRepo) GetTotalInstallment(ctx context.Context, customerID string, tenorMonth int) (float64, error) {
	return 0, nil
}

// ----------------------
// Token
// ----------------------

func (m *psqlMockRepo) CreateUserToken(ctx context.Context, payload models.Tokens) (err error) {
	return
}

func (m *psqlMockRepo) UpdateUserTokenByRefreshToken(ctx context.Context, refreshToken string, payload models.Tokens) (err error) {
	return
}

func (m *psqlMockRepo) FirstActiveRefreshToken(ctx context.Context, refreshToken string) (resp models.Tokens, err error) {
	return
}

func (m *psqlMockRepo) FirstActiveUserTokenByAccessToken(ctx context.Context, accessToken string) (resp models.Tokens, err error) {
	return
}

func (m *psqlMockRepo) FirstActiveUserTokenByUserId(ctx context.Context, userId string) (resp models.Tokens, err error) {
	return
}

func (m *psqlMockRepo) RevokeAllTokenByAccessTokenByUserId(ctx context.Context, userId string) (err error) {
	return
}
