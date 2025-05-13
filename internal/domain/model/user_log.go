package model

import (
	"time"
	"yourapp/pkg/core/model"
)

type UserLog struct {
	model.BaseModel
	UserID    uint64    `gorm:"not null;index"`
	Action    string    `gorm:"size:50;not null"`
	Details   string    `gorm:"type:text"`
	IPAddress string    `gorm:"size:45"`
	UserAgent string    `gorm:"size:255"`
	CreatedAt time.Time `gorm:"not null"`
	User      *User     `gorm:"foreignKey:UserID"`
}

const (
	ActionUserCreated = "user_created"
	ActionUserUpdated = "user_updated"
	ActionUserDeleted = "user_deleted"
)
