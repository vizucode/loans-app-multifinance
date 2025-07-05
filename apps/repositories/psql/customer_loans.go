package psql

import (
	"context"
	"errors"
	"multifinancetest/apps/models"

	"github.com/vizucode/gokit/logger"
	"gorm.io/gorm"
)

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
