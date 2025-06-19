package schema

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

// UserResponse represents the response schema for user info
type UserResponse struct {
	Data *UserInfo `json:"data,omitempty"`
}

// UserRecord represents a user record in datatable
type UserRecord struct {
	UserInfo
}

// UserDataTableResult represents the datatable result
type UserDataTableResult struct {
	Items []UserRecord `json:"items"`
	Total int64        `json:"total"`
}

// UserDatatableResponse represents the response schema for datatable
type UserDatatableResponse struct {
	Data UserDataTableResult `json:"data"`
}

// UserBatchRequest represents a batch request
type UserBatchRequest struct {
	Items []int `json:"items" validate:"required"`
}
