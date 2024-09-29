package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Product struct
type Product struct {
	gorm.Model
	ID          uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Title       string    `gorm:"not null" json:"title"`
	Description string    `gorm:"not null" json:"description"`
	Amount      int       `gorm:"not null" json:"amount"`
}
