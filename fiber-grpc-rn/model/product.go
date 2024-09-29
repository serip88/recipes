package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Product struct
type Product struct {
	gorm.Model
	ID          uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Title       string    `gorm:"type:varchar;not null" json:"title"`
	Description string    `gorm:"type:varchar;not null" json:"description"`
	Amount      float64   `gorm:"type:double precision;not null" json:"amount"`
}
