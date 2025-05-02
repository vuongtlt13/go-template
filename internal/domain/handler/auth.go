package handler

import (
	"connectrpc.com/connect"
	"context"
	"yourapp/internal/domain/service"
	"yourapp/pb/auth"
)

type AuthHandler struct {
	svc service.AuthService
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		svc: service.NewAuthService(),
	}
}

func (h *AuthHandler) Login(ctx context.Context, req *connect.Request[auth.LoginRequest]) (*connect.Response[auth.AuthResponse], error) {
	token, err := h.svc.Login(ctx, service.Credential{
		Email:    req.Msg.Email,
		Password: req.Msg.Password,
	})

	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&auth.AuthResponse{AccessToken: token}), nil
}

func (h *AuthHandler) Register(ctx context.Context, req *connect.Request[auth.RegisterRequest]) (*connect.Response[auth.AuthResponse], error) {
	//TODO implement me
	panic("implement me")
}
