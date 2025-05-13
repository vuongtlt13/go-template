package model

import (
	"yourapp/pkg/core/model"
)

type UserRole struct {
	model.BaseModel
	UserID uint64 `gorm:"not null;index:uix_user_role,unique"`
	RoleID uint64 `gorm:"not null;index:uix_user_role,unique"`

	User *User `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE" json:"-"`
	Role *Role `gorm:"foreignKey:RoleID;references:ID;constraint:OnDelete:CASCADE" json:"-"`
}

func (UserRole) TableName() string {
	return "user_roles"
}
