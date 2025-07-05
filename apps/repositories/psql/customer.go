package psql

import (
	"context"
	"errors"

	"multifinancetest/apps/models"

	"github.com/vizucode/gokit/logger"
	"gorm.io/gorm"
)

func (psql *psql) CreateCustomer(ctx context.Context, payload models.Customer) (err error) {

	if err := psql.db.Model(&models.Customer{}).WithContext(ctx).Create(&payload).Error; err != nil {
		logger.Log.Error(ctx, err)
		return err
	}

	return nil
}

func (psql *psql) FirstCustomerById(ctx context.Context, id string) (resp models.Customer, err error) {

	if err := psql.db.Model(&models.Customer{}).WithContext(ctx).Where("id = ?", id).First(&resp).Error; err != nil {
		logger.Log.Error(ctx, err)
		return resp, err
	}

	return resp, nil
}

func (psql *psql) FirstCustomerByEmail(ctx context.Context, email string) (resp models.Customer, err error) {
	if err := psql.db.Model(&models.Customer{}).WithContext(ctx).Where("email = ?", email).First(&resp).Error; err != nil {
		logger.Log.Error(ctx, err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return resp, nil
		}
		return resp, err
	}

	return resp, nil
}
