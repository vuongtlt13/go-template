package schema

import "yourapp/pkg/datatable"

// UserIDParam represents path parameter for user ID
type UserIDParam struct {
	ID uint `parse:"path:id" validate:"required,gt=0" description:"User ID"`
}

// GetUsersRequest represents query parameters for user endpoints
type GetUsersRequest struct {
	*datatable.RequestParams
	IsActive *bool `parse:"query:isActive" validate:"omitempty" description:"Filter by active status"`
	IsAdmin  *bool `parse:"query:isAdmin" validate:"omitempty" description:"Filter by admin status"`
}

func (r *GetUsersRequest) GetRequestParams() *datatable.RequestParams {
	return r.RequestParams
}

// UserBase represents shared properties in all create/update/read processes
type UserBase struct {
	Email    string `json:"email" validate:"required,email"`
	FullName string `json:"fullName" validate:"required,min=1,max=255"`
	IsActive bool   `json:"isActive" default:"true"`
	IsAdmin  bool   `json:"isAdmin" default:"false"`
}

// UserInDBBase represents properties in DB
type UserInDBBase struct {
	UserBase
	ID uint `json:"id" validate:"required,gt=0"`
}

// UserCreateRequest represents properties to receive via API on creation
type UserCreateRequest struct {
	UserBase
	Password string `json:"password" validate:"required,min=1,max=255"`
}

// UserUpdateRequest represents properties to receive via API on update
type UserUpdateRequest struct {
	UserBase
	Password *string `json:"password,omitempty" validate:"omitempty,min=1,max=255"`
}

// ProfileUpdateRequest represents properties to update user profile
type ProfileUpdateRequest struct {
	FullName *string `json:"fullName,omitempty" validate:"omitempty,min=1,max=255"`
	Password *string `json:"password,omitempty" validate:"omitempty,min=1,max=255"`
}

// UserInfo represents additional properties to return via API
type UserInfo struct {
	UserInDBBase
	UserRoles []string `json:"userRoles,omitempty"`
}

// UserRecord represents a user record in datatable
type UserRecord struct {
	UserInfo
}

// UserDataTableResult represents the datatable result
type UserDataTableResult = datatable.Result[interface{}]

// UserBatchRequest represents a batch request
type UserBatchRequest struct {
	Items []int `json:"items" validate:"required"`
}
