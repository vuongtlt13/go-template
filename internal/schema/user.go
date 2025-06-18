package schema

import (
	"yourapp/pkg/schema"
)

// UserBase represents shared properties in all create/update/read processes
type UserBase struct {
	Email    string `json:"email" validate:"required,email"`
	FullName string `json:"fullName" validate:"required,min=1,max=255"`
	IsActive bool   `json:"isActive" default:"true"`
	IsAdmin  bool   `json:"isAdmin" default:"false"`
}

// NewUserBase creates a new UserBase with validation
func NewUserBase(email, fullName string, isActive, isAdmin bool) (*UserBase, error) {
	user := &UserBase{
		Email:    email,
		FullName: fullName,
		IsActive: isActive,
		IsAdmin:  isAdmin,
	}
	if err := schema.Validate(user); err != nil {
		return nil, err
	}
	return user, nil
}

// UserInDBBase represents properties in DB
type UserInDBBase struct {
	UserBase
	ID uint `json:"id" validate:"required,gt=0"`
}

// NewUserInDBBase creates a new UserInDBBase with validation
func NewUserInDBBase(id uint, base *UserBase) (*UserInDBBase, error) {
	user := &UserInDBBase{
		UserBase: *base,
		ID:       id,
	}
	if err := schema.Validate(user); err != nil {
		return nil, err
	}
	return user, nil
}

// UserCreateRequest represents properties to receive via API on creation
type UserCreateRequest struct {
	UserBase
	Password string `json:"password" validate:"required,min=1,max=255"`
}

// NewUserCreateRequest creates a new UserCreateRequest with validation
func NewUserCreateRequest(email, fullName, password string, isActive, isAdmin bool) (*UserCreateRequest, error) {
	base, err := NewUserBase(email, fullName, isActive, isAdmin)
	if err != nil {
		return nil, err
	}

	user := &UserCreateRequest{
		UserBase: *base,
		Password: password,
	}
	if err := schema.Validate(user); err != nil {
		return nil, err
	}
	return user, nil
}

// UserUpdateRequest represents properties to receive via API on update
type UserUpdateRequest struct {
	UserBase
	Password *string `json:"password,omitempty" validate:"omitempty,min=1,max=255"`
}

// NewUserUpdateRequest creates a new UserUpdateRequest with validation
func NewUserUpdateRequest(email, fullName string, isActive, isAdmin bool, password *string) (*UserUpdateRequest, error) {
	base, err := NewUserBase(email, fullName, isActive, isAdmin)
	if err != nil {
		return nil, err
	}

	user := &UserUpdateRequest{
		UserBase: *base,
		Password: password,
	}
	if err := schema.Validate(user); err != nil {
		return nil, err
	}
	return user, nil
}

// ProfileUpdateRequest represents properties to update user profile
type ProfileUpdateRequest struct {
	FullName *string `json:"fullName,omitempty" validate:"omitempty,min=1,max=255"`
	Password *string `json:"password,omitempty" validate:"omitempty,min=1,max=255"`
}

// NewProfileUpdateRequest creates a new ProfileUpdateRequest with validation
func NewProfileUpdateRequest(fullName, password *string) (*ProfileUpdateRequest, error) {
	user := &ProfileUpdateRequest{
		FullName: fullName,
		Password: password,
	}
	if err := schema.Validate(user); err != nil {
		return nil, err
	}
	return user, nil
}

// UserInfo represents additional properties to return via API
type UserInfo struct {
	UserInDBBase
	UserRoles []string `json:"userRoles,omitempty"`
}

// NewUserInfo creates a new UserInfo with validation
func NewUserInfo(id uint, base *UserBase, roles []string) (*UserInfo, error) {
	dbBase, err := NewUserInDBBase(id, base)
	if err != nil {
		return nil, err
	}

	user := &UserInfo{
		UserInDBBase: *dbBase,
		UserRoles:    roles,
	}
	if err := schema.Validate(user); err != nil {
		return nil, err
	}
	return user, nil
}

// UserResponse represents the response schema for user info
type UserResponse struct {
	Data *UserInfo `json:"data,omitempty"`
}

// NewUserResponse creates a new UserResponse with validation
func NewUserResponse(data *UserInfo) (*UserResponse, error) {
	response := &UserResponse{
		Data: data,
	}
	if err := schema.Validate(response); err != nil {
		return nil, err
	}
	return response, nil
}

// UserRecord represents a user record in datatable
type UserRecord struct {
	UserInfo
}

// NewUserRecord creates a new UserRecord with validation
func NewUserRecord(info *UserInfo) (*UserRecord, error) {
	record := &UserRecord{
		UserInfo: *info,
	}
	if err := schema.Validate(record); err != nil {
		return nil, err
	}
	return record, nil
}

// UserDataTableResult represents the datatable result
type UserDataTableResult struct {
	Items []UserRecord `json:"items"`
	Total int64        `json:"total"`
}

// NewUserDataTableResult creates a new UserDataTableResult with validation
func NewUserDataTableResult(items []UserRecord, total int64) (*UserDataTableResult, error) {
	result := &UserDataTableResult{
		Items: items,
		Total: total,
	}
	if err := schema.Validate(result); err != nil {
		return nil, err
	}
	return result, nil
}

// UserDatatableResponse represents the response schema for datatable
type UserDatatableResponse struct {
	Data UserDataTableResult `json:"data"`
}

// NewUserDatatableResponse creates a new UserDatatableResponse with validation
func NewUserDatatableResponse(data UserDataTableResult) (*UserDatatableResponse, error) {
	response := &UserDatatableResponse{
		Data: data,
	}
	if err := schema.Validate(response); err != nil {
		return nil, err
	}
	return response, nil
}

// UserBatchRequest represents a batch request
type UserBatchRequest struct {
	Items []int `json:"items" validate:"required"`
}

// NewUserBatchRequest creates a new UserBatchRequest with validation
func NewUserBatchRequest(items []int) (*UserBatchRequest, error) {
	request := &UserBatchRequest{
		Items: items,
	}
	if err := schema.Validate(request); err != nil {
		return nil, err
	}
	return request, nil
}
