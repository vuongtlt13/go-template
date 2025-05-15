package model

import (
	"time"

	"gorm.io/gorm"
)

const (
	RoleUser       = "user"
	RoleAdmin      = "admin"
	RoleSuperAdmin = "superadmin"
)

// Role represents a user role
type Role struct {
	ID          uint64         `gorm:"primarykey" json:"id"`
	Code        string         `gorm:"size:50;not null;uniqueIndex" json:"code"`
	Name        string         `gorm:"size:100;not null" json:"name"`
	Description string         `gorm:"size:255" json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	Users       []User         `gorm:"many2many:user_roles;" json:"users,omitempty"`
	Permissions []Permission   `gorm:"many2many:role_permissions;" json:"permissions,omitempty"`
}
