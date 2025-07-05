package psql

import (
	"context"
	"errors"
	"multifinancetest/apps/models"
	"multifinancetest/helpers/constants/rpcstd"
	"net/http"

	"github.com/google/uuid"
	"github.com/vizucode/gokit/logger"
	"github.com/vizucode/gokit/utils/errorkit"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (psql *psql) CreateTrxCustomerLoan(ctx context.Context, customerID string, tenorMonth int, req models.CustomerLoans) (err error) {
	return psql.db.Transaction(func(tx *gorm.DB) error {
		var customerLimit models.CustomerTenors
		err := tx.Table("customer_tenors ct").
			WithContext(ctx).
			Clauses(clause.Locking{Strength: "UPDATE"}).
			Joins("JOIN tenors t ON t.id = ct.tenor_id").
			Where("ct.customer_id = ? AND t.total_month = ?", customerID, tenorMonth).
			First(&customerLimit).Error
		if err != nil {
			return errorkit.NewErrorStd(http.StatusBadRequest, rpcstd.ABORTED, "out of limit tenor")
		}

		var totalUsed float64
		err = tx.Model(&models.CustomerLoans{}).
			WithContext(ctx).
			Select("COALESCE(SUM(otr), 0)").
			Where("customer_id = ? AND total_month = ?", customerID, tenorMonth).
			Scan(&totalUsed).Error
		if err != nil {
			logger.Log.Error(ctx, err)
			return errorkit.NewErrorStd(http.StatusInternalServerError, rpcstd.ABORTED, "failed to count limit remaining tenor month loan")
		}

		remaining := float64(customerLimit.LimitLoanAmount) - totalUsed

		if req.Otr > remaining {
			return errorkit.NewErrorStd(http.StatusBadRequest, rpcstd.ABORTED, "out of remaining loan")
		}

		adminFee := 0.1
		interestRate := 0.5

		interestAmount := req.Otr * interestRate * float64(tenorMonth)
		adminFeeAmount := req.Otr * adminFee
		totalLoan := req.Otr + interestAmount + adminFeeAmount
		monthlyInstallment := totalLoan / float64(tenorMonth)

		err = tx.Create(&models.CustomerLoans{
			ID:                     uuid.NewString(),
			CustomerID:             customerID,
			AssetName:              req.AssetName,
			Otr:                    req.Otr,
			TotalMonth:             tenorMonth,
			InterestRate:           interestRate,
			AdminFeeAmount:         adminFeeAmount,
			InterestAmount:         interestAmount,
			InstallmentAmount:      monthlyInstallment,
			TotalInstallmentAmount: totalLoan,
		}).Error
		if err != nil {
			logger.Log.Error(ctx, err)
			return errorkit.NewErrorStd(http.StatusInternalServerError, rpcstd.ABORTED, "failed to create customer loan")
		}

		return nil
	})
}

func (psql *psql) CreateCustomerLoan(ctx context.Context, payload models.CustomerLoans) (err error) {

	err = psql.db.Model(&models.CustomerLoans{}).WithContext(ctx).Create(&payload).Error
	if err != nil {
		return err
	}

	return nil
}

func (psql *psql) CountLimitRemainingTenorMonthLoan(ctx context.Context, customerID string, tenorMonth int) (resp int, err error) {

	var total float64

	Row := psql.db.Model(&models.CustomerLoans{}).
		WithContext(ctx).
		Select("COALESCE(SUM(otr), 0)").
		Where("customer_id = ? AND total_month = ?", customerID, tenorMonth).
		Row()

	err = Row.Scan(&total)
	if err != nil {
		return 0, err
	}

	customerLimit, err := psql.GetCustomerLimitTenor(ctx, customerID, tenorMonth)
	if err != nil {
		return 0, err
	}

	if total > 0 {
		total = customerLimit.LimitLoanAmount - total
	} else {
		total = float64(customerLimit.LimitLoanAmount)
	}

	return int(total), nil
}

func (psql *psql) GetCustomerLoans(ctx context.Context, customerId string, filter models.Filter) (resp []models.CustomerLoans, err error) {

	tx := psql.db.Debug().WithContext(ctx).Model(&models.CustomerLoans{}).Where("customer_id = ?", customerId).Order("created_at DESC")

	if filter.Where["search"].(string) != "" {
		tx = tx.Where("asset_name = ?", filter.Where["search"].(string))
	}

	if filter.Limit > 0 {
		tx = tx.Limit(filter.Limit)
	}

	if filter.Page > 0 {
		tx = tx.Offset((filter.Page - 1) * filter.Limit)
	}

	err = tx.Find(&resp).Error
	if err != nil {
		logger.Log.Error(ctx, err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return resp, nil
		}
		return resp, err
	}

	return resp, nil
}
