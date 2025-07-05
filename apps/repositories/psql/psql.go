package psql

import "gorm.io/gorm"

type psql struct {
	db *gorm.DB
}

func NewPsql(
	db *gorm.DB,
) *psql {
	return &psql{
		db: db,
	}
}
