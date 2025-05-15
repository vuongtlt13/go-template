package server

import (
	"context"
	"fmt"

	"yourapp/internal/handler"
	"yourapp/pb/auth/authconnect"
	"yourapp/pkg/config"
	"yourapp/pkg/server"

	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

// UserServer represents the user HTTP server
type UserServer struct {
	*server.BaseServer
}

// NewUserServer creates a new user server instance
func NewUserServer(cfg *config.Config) *UserServer {
	return &UserServer{
		BaseServer: server.NewBaseServer(cfg, "User"),
	}
}

// Additional user-specific methods

// Start starts the server
func (s *UserServer) Start() error {
	// Initialize handlers
	authHandler := handler.NewAuthHandler()

	// Create Connect service handlers
	authPath, authConnectHandler := authconnect.NewAuthServiceHandler(authHandler)

	// Mount Connect services
	app := s.GetApp()
	app.All(authPath+"*", adaptor.HTTPHandler(authConnectHandler))

	// Start server
	return app.Listen(fmt.Sprintf(":%d", s.GetConfig().Server.UserPort))
}

// Shutdown gracefully shuts down the server
func (s *UserServer) Shutdown(ctx context.Context) error {
	return s.BaseServer.Shutdown(ctx)
}
