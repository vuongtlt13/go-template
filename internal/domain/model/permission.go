package model

import (
	"yourapp/internal/core/model"
)

type Permission struct {
	model.BaseModel
	Name   string  `gorm:"size:255;unique;not null"`
	Method *string `gorm:"size:10"`
	URL    *string `gorm:"size:255"`
	Roles  []*Role `gorm:"many2many:role_permissions;joinForeignKey:PermissionID"`
}
