package model

import (
	"yourapp/internal/core/model"
)

type Role struct {
	model.BaseModel
	Name        string        `gorm:"size:255;not null"`
	Code        string        `gorm:"size:31;unique;not null"`
	Users       []*User       `gorm:"many2many:user_roles;joinForeignKey:RoleID;joinReferences:UserID"`
	Permissions []*Permission `gorm:"many2many:role_permissions;joinForeignKey:RoleID;joinReferences:PermissionID"`
}
