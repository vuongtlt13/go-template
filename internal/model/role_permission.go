package model

type RolePermission struct {
	RoleID       uint64 `gorm:"not null;index:uix_role_permission,unique"`
	PermissionID uint64 `gorm:"not null;index:uix_role_permission,unique"`
}

func (RolePermission) TableName() string {
	return "role_permissions"
}
