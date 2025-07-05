package psql

import (
	"context"
	"multifinancetest/apps/models"
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
