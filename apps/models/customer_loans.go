package models

import "time"

type CustomerLoans struct {
	ID                     string    `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	CustomerID             string    `gorm:"type:uuid;column:customer_id"`
	AssetName              string    `gorm:"column:asset_name"`
	Otr                    float64   `gorm:"column:otr"`
	TotalMonth             int       `gorm:"column:total_month"`
	InterestRate           float64   `gorm:"column:interest_rate"`
	AdminFeeAmount         float64   `gorm:"column:admin_fee_amount"`
	InterestAmount         float64   `gorm:"column:interest_amount"`
	InstallmentAmount      float64   `gorm:"column:installment_amount"`
	TotalInstallmentAmount float64   `gorm:"column:total_installment_amount"`
	CreatedAt              time.Time `gorm:"column:created_at"`
	UpdatedAt              time.Time `gorm:"column:updated_at"`
}
