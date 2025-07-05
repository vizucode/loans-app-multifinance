package psql

import (
	"context"
	"multifinancetest/apps/models"

	"github.com/vizucode/gokit/logger"
)

func (psql *psql) CreateCustomerTenor(ctx context.Context, payload models.CustomerTenors) (err error) {

	err = psql.db.Model(&models.CustomerTenors{}).WithContext(ctx).Create(&payload).Error
	if err != nil {
		logger.Log.Error(ctx, err)
		return err
	}

	return nil
}

func (psql *psql) GetAllCustomerTenor(ctx context.Context, customerID string) (resp []models.CustomerTenors, err error) {

	err = psql.db.Model(&models.CustomerTenors{}).WithContext(ctx).Where("customer_id = ?", customerID).Preload("Tenor").Find(&resp).Error
	if err != nil {
		logger.Log.Error(ctx, err)
		return resp, err
	}

	return resp, nil
}
