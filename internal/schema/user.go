package schema

// UserIDParam represents path parameter for user ID
type UserIDParam struct {
	ID uint `parse:"path:id" validate:"required,gt=0" description:"User ID"`
}

// PaginationQuery represents common pagination query parameters
type PaginationQuery struct {
	Page   int    `parse:"query:page" validate:"omitempty,min=1" default:"1" description:"Page number"`
	Limit  int    `parse:"query:limit" validate:"omitempty,min=1,max=100" default:"10" description:"Number of items per page"`
	Search string `parse:"query:search" validate:"omitempty,max=255" description:"Search term"`
}

// UserQuery represents query parameters for user endpoints
type UserQuery struct {
	PaginationQuery
	IsActive *bool `parse:"query:isActive" validate:"omitempty" description:"Filter by active status"`
	IsAdmin  *bool `parse:"query:isAdmin" validate:"omitempty" description:"Filter by admin status"`
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
