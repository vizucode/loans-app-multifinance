package loan

import (
	"context"
	"multifinancetest/apps/domain"
	"multifinancetest/apps/middlewares/security"
	"multifinancetest/apps/models"
	"multifinancetest/helpers/constants/rpcstd"
	"net/http"
	"time"

	"github.com/vizucode/gokit/logger"
	"github.com/vizucode/gokit/utils/errorkit"
)

func (uc *loan) GetLimitLoans(ctx context.Context) (resp []domain.ResponseLimitLoans, err error) {

	currentUser, found := security.ExtractUserContext(ctx)
	if !found {
		return resp, errorkit.NewErrorStd(http.StatusUnauthorized, rpcstd.ABORTED, "user not found")
	}

	respLimitLoans, err := uc.db.GetAllCustomerTenor(ctx, currentUser.Id)
	if err != nil {
		logger.Log.Error(ctx, err)
		return resp, err
	}

	for _, limitLoan := range respLimitLoans {

		remainingLoanAmount, err2 := uc.db.CountLimitRemainingTenorMonthLoan(ctx, currentUser.Id, limitLoan.Tenor.TotalMonth)
		if err2 != nil {
			logger.Log.Error(ctx, err2)
			return resp, err2
		}

		resp = append(resp, domain.ResponseLimitLoans{
			LoanId:              limitLoan.ID,
			LoanMonth:           limitLoan.Tenor.TotalMonth,
			TotalLoanAmount:     limitLoan.LimitLoanAmount,
			RemainingLoanAmount: float64(remainingLoanAmount),
		})
	}

	return resp, err
}

func (uc *loan) GetHistoryLoans(ctx context.Context, filter domain.Filter) (resp []domain.ResponseHistoryLoans, err error) {

	currentUser, found := security.ExtractUserContext(ctx)
	if !found {
		return resp, errorkit.NewErrorStd(http.StatusUnauthorized, rpcstd.ABORTED, "user not found")
	}

	resultHistoryPlan, err := uc.db.GetCustomerLoans(ctx, currentUser.Id, models.Filter{
		Limit: filter.Limit,
		Page:  filter.Page,
		Where: filter.Where,
	})
	if err != nil {
		logger.Log.Error(ctx, err)
		return resp, err
	}

	for _, historyPlan := range resultHistoryPlan {
		resp = append(resp, domain.ResponseHistoryLoans{
			LoanId:             historyPlan.ID,
			LoanMonth:          historyPlan.TotalMonth,
			Otr:                historyPlan.Otr,
			AssetName:          historyPlan.AssetName,
			MonthlyInstallment: historyPlan.InstallmentAmount,
			TotalInstallment:   historyPlan.TotalInstallmentAmount,
			CreatedAt:          historyPlan.CreatedAt.UTC().Format(time.RFC3339),
		})
	}

	return resp, nil
}
