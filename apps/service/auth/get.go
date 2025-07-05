package auth

import (
	"context"
	"multifinancetest/apps/domain"
	"multifinancetest/apps/middlewares/security"
	"multifinancetest/helpers/constants/rpcstd"
	"net/http"

	"github.com/vizucode/gokit/logger"
	"github.com/vizucode/gokit/utils/errorkit"
)

func (uc *auth) FirstCustomer(ctx context.Context) (resp domain.ResponseCustomer, err error) {

	currentUser, found := security.ExtractUserContext(ctx)
	if !found {
		return resp, errorkit.NewErrorStd(http.StatusUnauthorized, rpcstd.ABORTED, "user not found")
	}

	resultUCustomer, err := uc.db.FirstCustomerById(ctx, currentUser.Id)
	if err != nil {
		logger.Log.Error(ctx, err)
		return resp, err
	}

	resp = domain.ResponseCustomer{
		Email:                    resultUCustomer.Email,
		FullName:                 resultUCustomer.FullName,
		LegalName:                resultUCustomer.LegalName,
		DateBirth:                resultUCustomer.DateBirth.Format("2006-01-02"),
		BornAt:                   resultUCustomer.BornAt,
		Salary:                   resultUCustomer.Salary,
		NationalIdentityNumber:   resultUCustomer.NationalIdentityNumber,
		NationalIdentityImageURL: resultUCustomer.NationalIdentityImageUrl,
		SelfieImageURL:           resultUCustomer.SelfieImageUrl,
	}

	return resp, nil
}


