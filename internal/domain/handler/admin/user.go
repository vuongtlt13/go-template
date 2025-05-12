package admin

import (
	"context"
	"yourapp/internal/domain/repository"
	"yourapp/internal/domain/service"
	"yourapp/pb/admin"

	"connectrpc.com/connect"
	"gorm.io/gorm"
)

// UserHandler implements the Connect gRPC interface for admin user service
// It should satisfy adminconnect.UserServiceHandler

type UserHandler struct {
	userService service.UserService
}

// NewUserHandler creates a new user handler using singleton instances
func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{
		userService: service.NewUserService(db, repository.NewUserRepository()),
	}
}

// CreateUser creates a new user
func (h *UserHandler) CreateUser(ctx context.Context, req *connect.Request[admin.CreateUserRequest]) (*connect.Response[admin.UserResponse], error) {
	// TODO: Implement logic using h.svc
	return nil, connect.NewError(connect.CodeUnimplemented, nil)
}

// GetUser retrieves a user by ID
func (h *UserHandler) GetUser(ctx context.Context, req *connect.Request[admin.GetUserRequest]) (*connect.Response[admin.UserResponse], error) {
	// TODO: Implement logic using h.svc
	return nil, connect.NewError(connect.CodeUnimplemented, nil)
}

// UpdateUser updates a user
func (h *UserHandler) UpdateUser(ctx context.Context, req *connect.Request[admin.UpdateUserRequest]) (*connect.Response[admin.UserResponse], error) {
	// TODO: Implement logic using h.svc
	return nil, connect.NewError(connect.CodeUnimplemented, nil)
}

// DeleteUser deletes a user
func (h *UserHandler) DeleteUser(ctx context.Context, req *connect.Request[admin.DeleteUserRequest]) (*connect.Response[admin.DeleteUserResponse], error) {
	// TODO: Implement logic using h.svc
	return nil, connect.NewError(connect.CodeUnimplemented, nil)
}

// ListUsers retrieves all users
func (h *UserHandler) ListUsers(ctx context.Context, req *connect.Request[admin.ListUsersRequest]) (*connect.Response[admin.ListUsersResponse], error) {
	// TODO: Implement logic using h.svc
	return nil, connect.NewError(connect.CodeUnimplemented, nil)
}
