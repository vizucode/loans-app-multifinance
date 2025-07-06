package loan

import (
	"multifinancetest/apps/domain"
	"multifinancetest/apps/repositories/psqlmock"
	contextkeys "multifinancetest/helpers/constants/context_keys"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/net/context"
)

func TestCreateLoan(t *testing.T) {
	psqlMocked := psqlmock.NewPsqlMockRepo()

	ctx := context.WithValue(context.Background(), contextkeys.UserContext, domain.UserContext{
		Id:       "123",
		FullName: "sandi utomo",
		Exp:      1000000000000,
	})

	psqlMocked.On("CreateTrxCustomerLoan", ctx, mock.Anything, mock.Anything, mock.Anything).
		Return(nil)

	loanSvc := NewLoan(
		psqlMocked,
		validator.New(),
	)

	resultLoan, err := loanSvc.CreateLoan(ctx, domain.RequestLoans{
		Otr:         1000000,
		AssetName:   "BUGATI CHIRON",
		PickedTenor: 12,
	})
	if err != nil {
		assert.Error(t, err)
	}

	assert.NotEmpty(t, resultLoan)
}
