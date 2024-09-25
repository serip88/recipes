package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User struct
type User struct {
	gorm.Model
	ID       uuid.UUID `gorm:"primaryKey;type:varchar"`
	Username string    `gorm:"uniqueIndex;type:varchar;not null" json:"username"`
	Email    string    `gorm:"uniqueIndex;type:varchar;not null" json:"email"`
	Password string    `gorm:"type:varchar;not null" json:"password"`
	Names    string    `gorm:"type:varchar;not null" json:"names"`
}
