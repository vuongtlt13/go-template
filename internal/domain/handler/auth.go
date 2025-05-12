package handler

import (
	"context"
	"yourapp/internal/domain/repository"
	"yourapp/internal/domain/service"
	"yourapp/pb/auth"
	authpkg "yourapp/pkg/auth"
	"yourapp/pkg/database"

	"connectrpc.com/connect"
)

type AuthHandler struct {
	authService service.AuthService
}

// NewAuthHandler creates a new auth handler using singleton instances
func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		authService: service.NewAuthService(
			database.GetDatabase(),
			repository.NewUserRepository(),
			authpkg.GetJWTManager(),
		),
	}
}

// Login handles user login
func (h *AuthHandler) Login(ctx context.Context, req *connect.Request[auth.LoginRequest]) (*connect.Response[auth.AuthResponse], error) {
	token, err := h.authService.Login(ctx, service.Credential{
		Email:    req.Msg.Email,
		Password: req.Msg.Password,
	})

	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&auth.AuthResponse{AccessToken: token}), nil
}

// Register handles user registration
func (h *AuthHandler) Register(ctx context.Context, req *connect.Request[auth.RegisterRequest]) (*connect.Response[auth.AuthResponse], error) {
	// TODO: Implement registration logic using h.authService
	return nil, connect.NewError(connect.CodeUnimplemented, nil)
}
