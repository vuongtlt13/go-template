package model

import (
	"time"

	"gorm.io/gorm"
)

// Permission represents an API permission
type Permission struct {
	ID          uint64         `gorm:"primarykey" json:"id"`
	Code        string         `gorm:"size:100;not null;uniqueIndex" json:"code"`
	Name        string         `gorm:"size:100;not null" json:"name"`
	Description string         `gorm:"size:255" json:"description"`
	Service     string         `gorm:"size:50;not null" json:"service"` // e.g., "user", "role", "permission"
	Method      string         `gorm:"size:50;not null" json:"method"`  // e.g., "create", "read", "update", "delete"
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	Roles       []Role         `gorm:"many2many:role_permissions;" json:"roles,omitempty"`
}
