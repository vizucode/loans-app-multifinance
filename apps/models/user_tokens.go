package models

import (
	"time"
)

type Tokens struct {
	ID                    string    `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	CustomerId            string    `gorm:"column:customer_id;type:binary(16)"`
	AccessToken           string    `gorm:"column:access_token"`
	RefreshToken          string    `gorm:"column:refresh_token"`
	AccessTokenRevoked    bool      `gorm:"column:access_token_revoked"`
	RefreshTokenRevoked   bool      `gorm:"column:refresh_token_revoked"`
	AccessTokenExpiredAt  time.Time `gorm:"column:access_token_expired_at"`
	RefreshTokenExpiredAt time.Time `gorm:"column:refresh_token_expired_at"`
	CreatedAt             time.Time `gorm:"type:timestamptz;autoCreateTime" json:"created_at"`
	UpdatedAt             time.Time `gorm:"type:timestamptz;autoUpdateTime" json:"updated_at"`
}
