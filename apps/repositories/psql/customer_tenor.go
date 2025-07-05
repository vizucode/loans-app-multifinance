package psql

import (
	"context"
	"errors"
	"multifinancetest/apps/models"

	"github.com/vizucode/gokit/logger"
	"gorm.io/gorm"
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

func (psql *psql) GetCustomerLimitTenor(ctx context.Context, customerID string, tenorMonth int) (resp models.CustomerTenors, err error) {

	tx := psql.db.Table("customer_tenors ct").
		WithContext(ctx).
		Where("ct.customer_id = ?", customerID)

	if tenorMonth > 0 {
		tx = tx.Joins("JOIN tenors t ON t.id = ct.tenor_id").
			Where("t.total_month = ?", tenorMonth)
	}

	err = tx.First(&resp).Error
	if err != nil {
		logger.Log.Error(ctx, err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return resp, nil
		}
		return resp, err
	}

	return resp, nil
}
