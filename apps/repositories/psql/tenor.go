package psql

import (
	"context"
	"multifinancetest/apps/models"

	"github.com/vizucode/gokit/logger"
)

func (psql *psql) GetAllTenor(ctx context.Context) (resp []models.Tenor, err error) {

	err = psql.db.Model(&models.Tenor{}).WithContext(ctx).Find(&resp).Error
	if err != nil {
		logger.Log.Error(ctx, err)
		return resp, err
	}

	return resp, nil
}
