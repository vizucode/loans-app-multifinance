package models

import (
	"time"
)

type Customer struct {
	ID                       string    `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Email                    string    `gorm:"column:email"`
	Password                 string    `gorm:"column:password"`
	FullName                 string    `gorm:"column:full_name"`
	LegalName                string    `gorm:"column:legal_name"`
	DateBirth                time.Time `gorm:"column:date_birth"`
	BornAt                   string    `gorm:"column:born_at"`
	Salary                   float64   `gorm:"column:salary"`
	NationalIdentityNumber   string    `gorm:"column:national_identity_number"`
	NationalIdentityImageUrl string    `gorm:"column:national_identity_image_url"`
	SelfieImageUrl           string    `gorm:"column:selfie_image_url"`
	CreatedAt                time.Time `gorm:"column:created_at"`
	UpdatedAt                time.Time `gorm:"column:updated_at"`
}
