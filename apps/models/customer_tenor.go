package models

import (
	"time"
)

type CustomerTenors struct {
	ID              string    `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	CustomerID      string    `gorm:"type:uuid;column:customer_id"`
	TenorID         string    `gorm:"type:uuid;column:tenor_id"`
	LimitLoanAmount float64   `gorm:"column:limit_loan_amount"`
	CreatedAt       time.Time `gorm:"column:created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at"`

	Tenor    Tenor    `gorm:"foreignKey:TenorID;references:ID"`
	Customer Customer `gorm:"foreignKey:CustomerID;references:ID"`
}
