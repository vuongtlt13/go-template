package service

import "errors"

var (
	ErrUserNotFound            = errors.New("user not found")
	ErrDuplicateEmail          = errors.New("duplicate email")
	ErrRoleNotFound            = errors.New("role not found")
	ErrDuplicateRoleCode       = errors.New("duplicate role code")
	ErrPermissionNotFound      = errors.New("permission not found")
	ErrDuplicatePermissionCode = errors.New("duplicate permission code")
)
