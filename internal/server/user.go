package server

import (
	"context"
	"fmt"
	"yourapp/internal/routes"
	"yourapp/pkg/config"
	"yourapp/pkg/logger"
	"yourapp/pkg/server"
)

// User represents the user HTTP server
type UserServer struct {
	*server.BaseServer
	logger logger.Logger
}

// NewUserServer creates a new user server instance
func NewUserServer(cfg *config.Config, logger logger.Logger) *UserServer {
	return &UserServer{
		BaseServer: server.NewBaseServer(cfg, "User"),
		logger:     logger,
	}
}

// Start starts the server
func (s *UserServer) Start() error {
	app := s.GetApp()

	// Register user routes (bao gồm cả common routes)
	routes.NewUserRouter().Register(app)

	// Auto-fiber OpenAPI docs and Swagger UI
	app.ServeDocs("/docs")
	app.ServeSwaggerUI("/swagger", "/docs")

	// Start server
	addr := fmt.Sprintf(":%d", s.GetConfig().Server.Port)
	s.logger.Infof("Starting user server at %s", addr)
	return app.Listen(addr)
}

// Shutdown gracefully shuts down the server
func (s *UserServer) Shutdown(ctx context.Context) error {
	return s.BaseServer.Shutdown(ctx)
}
