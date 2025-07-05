package models

import (
	"time"

	"github.com/google/uuid"
)

type Tenor struct {
	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	TotalMonth int       `gorm:"column:total_month"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at"`
}
