package loan

import (
	"errors"
	"multifinancetest/apps/domain"
	"multifinancetest/apps/repositories/psqlmock"
	contextkeys "multifinancetest/helpers/constants/context_keys"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/net/context"
)

func setupTest() (context.Context, *validator.Validate, domain.UserContext) {
	validate := validator.New()
	ctx := context.WithValue(context.Background(), contextkeys.UserContext, domain.UserContext{
		Id:       "123",
		FullName: "Sandi Utomo",
		Exp:      1000000000000,
	})
	return ctx, validate, ctx.Value(contextkeys.UserContext).(domain.UserContext)
}

func TestCreateLoan(t *testing.T) {
	ctx, validate, _ := setupTest()

	t.Run("success create loan", func(t *testing.T) {
		psqlMock := psqlmock.NewPsqlMockRepo()
		loanSvc := NewLoan(psqlMock, validate)

		req := domain.RequestLoans{
			Otr:         1000000,
			AssetName:   "BUGATTI CHIRON",
			PickedTenor: 12,
		}

		psqlMock.On("CreateTrxCustomerLoan", ctx, "123", req.PickedTenor, mock.Anything).Return(nil).Once()

		result, err := loanSvc.CreateLoan(ctx, req)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, "BUGATTI CHIRON", result.AssetName)
	})

	t.Run("validation fails and skips DB call", func(t *testing.T) {
		psqlMock := psqlmock.NewPsqlMockRepo()
		loanSvc := NewLoan(psqlMock, validate)

		req := domain.RequestLoans{
			Otr:         500000,
			AssetName:   "AVANZA 2023",
			PickedTenor: 0, // invalid
		}

		_, err := loanSvc.CreateLoan(ctx, req)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "PickedTenor")
		psqlMock.AssertNotCalled(t, "CreateTrxCustomerLoan")
	})

	t.Run("database error after valid request", func(t *testing.T) {
		psqlMock := psqlmock.NewPsqlMockRepo()
		loanSvc := NewLoan(psqlMock, validate)

		req := domain.RequestLoans{
			Otr:         1000000,
			AssetName:   "TOYOTA SUPRA",
			PickedTenor: 24,
		}

		psqlMock.On("CreateTrxCustomerLoan", ctx, "123", req.PickedTenor, mock.Anything).
			Return(errors.New("database connection failed")).Once()

		_, err := loanSvc.CreateLoan(ctx, req)

		assert.Error(t, err)
		assert.EqualError(t, err, "database connection failed")
	})
}

func TestRequestLoanValidation(t *testing.T) {
	validate := validator.New()

	t.Run("invalid struct should return validation errors", func(t *testing.T) {
		req := domain.RequestLoans{
			Otr:         0,
			AssetName:   "",
			PickedTenor: 0,
		}

		err := validate.Struct(req)
		assert.Error(t, err)

		validationErrs, ok := err.(validator.ValidationErrors)
		assert.True(t, ok)
		assert.Greater(t, len(validationErrs), 0)
	})

	t.Run("valid struct should pass validation", func(t *testing.T) {
		req := domain.RequestLoans{
			Otr:         1000000,
			AssetName:   "Valid Asset",
			PickedTenor: 12,
		}

		err := validate.Struct(req)
		assert.NoError(t, err)
	})
}
