package model

import (
	"time"

	"gorm.io/gorm"
)

// UserProfile represents a user's profile information
type UserProfile struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	UserID    uint           `gorm:"not null;uniqueIndex" json:"user_id"`
	FirstName string         `gorm:"size:50" json:"first_name"`
	LastName  string         `gorm:"size:50" json:"last_name"`
	Phone     string         `gorm:"size:20" json:"phone"`
	Address   string         `gorm:"size:255" json:"address"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
