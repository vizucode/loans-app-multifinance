package loan

import (
	"context"
	"multifinancetest/apps/domain"
	"multifinancetest/apps/middlewares/security"
	"multifinancetest/apps/models"
	"multifinancetest/helpers/constants/rpcstd"
	"net/http"

	"github.com/google/uuid"
	"github.com/vizucode/gokit/logger"
	"github.com/vizucode/gokit/utils/errorkit"
)

func (uc *loan) CreateLoan(ctx context.Context, req domain.RequestLoans) (resp domain.ResponseLoans, err error) {
	var (
		adminFee     float64 = 0.1
		interestRate float64 = 0.5
	)
	currentUser, found := security.ExtractUserContext(ctx)
	if !found {
		return resp, errorkit.NewErrorStd(http.StatusUnauthorized, rpcstd.ABORTED, "user not found")
	}

	err = uc.validateLoan(ctx, req, currentUser.Id)
	if err != nil {
		logger.Log.Error(ctx, err)
		return resp, err
	}

	interestAmount := req.Otr * interestRate * float64(req.PickedTenor)
	adminFeeAmount := (adminFee * req.Otr)
	totalLoan := req.Otr + interestAmount + adminFeeAmount
	monthlyInstallment := totalLoan / float64(req.PickedTenor)

	loanModel := models.CustomerLoans{
		ID:                     uuid.NewString(),
		CustomerID:             currentUser.Id,
		AssetName:              req.AssetName,
		Otr:                    req.Otr,
		TotalMonth:             req.PickedTenor,
		InterestRate:           interestRate,
		AdminFeeAmount:         adminFeeAmount,
		InterestAmount:         interestAmount,
		InstallmentAmount:      monthlyInstallment,
		TotalInstallmentAmount: totalLoan,
	}

	err = uc.db.CreateCustomerLoan(ctx, loanModel)
	if err != nil {
		logger.Log.Error(ctx, err)
		return resp, err
	}

	resp.AdminFee = adminFeeAmount
	resp.AssetName = req.AssetName
	resp.InterestFee = interestAmount
	resp.InterestRate = interestRate
	resp.Otr = req.Otr
	resp.PickedTenor = req.PickedTenor
	resp.TotalMonthlyInstalement = monthlyInstallment
	resp.TotalInstallmentAmount = totalLoan

	return resp, err
}

func (uc *loan) validateLoan(ctx context.Context, req domain.RequestLoans, customerID string) (err error) {

	err = uc.validator.StructCtx(ctx, req)
	if err != nil {
		return err
	}

	customerLimit, err := uc.db.GetCustomerLimitTenor(ctx, customerID, req.PickedTenor)
	if err != nil {
		logger.Log.Error(ctx, err)
		return err
	}

	// validate limit OTR
	if float64(customerLimit.LimitLoanAmount) < req.Otr {
		return errorkit.NewErrorStd(http.StatusBadRequest, rpcstd.ABORTED, "out of limit loan")
	}

	remainingLoan, err := uc.db.CountLimitRemainingTenorMonthLoan(ctx, customerID, req.PickedTenor)
	if err != nil {
		logger.Log.Error(ctx, err)
		return err
	}

	// validate remaining loan
	if float64(remainingLoan) < req.Otr {
		return errorkit.NewErrorStd(http.StatusBadRequest, rpcstd.ABORTED, "out of remaining loan")
	}

	return nil
}
