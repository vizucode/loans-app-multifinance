package loan

import (
	"context"
	"multifinancetest/apps/domain"
	"multifinancetest/apps/middlewares/security"
	"multifinancetest/helpers/constants/rpcstd"
	"net/http"

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
