package server

import (
	"context"
	"fmt"

	"yourapp/internal/routes"
	"yourapp/pkg/config"
	"yourapp/pkg/logger"
	"yourapp/pkg/server"
)

// Admin represents the admin HTTP server
type AdminServer struct {
	*server.BaseServer
	logger logger.Logger
}

// NewAdminServer creates a new admin server instance
func NewAdminServer(cfg *config.Config, logger logger.Logger) *AdminServer {
	return &AdminServer{
		BaseServer: server.NewBaseServer(cfg, "Admin"),
		logger:     logger,
	}
}

// Start starts the server
func (s *AdminServer) Start() error {
	app := s.GetApp()

	// Initialize and register admin routes
	adminRouter := routes.NewAdminRouter()
	adminRouter.Register(app)

	// Auto-fiber OpenAPI docs and Swagger UI
	app.ServeDocs("/docs")
	app.ServeSwaggerUI("/swagger", "/docs")

	// Start server
	addr := fmt.Sprintf(":%d", s.GetConfig().Server.Port)
	s.logger.Info("Starting admin server", "addr", addr)
	return app.Listen(addr)
}

// Shutdown gracefully shuts down the server
func (s *AdminServer) Shutdown(ctx context.Context) error {
	return s.BaseServer.Shutdown(ctx)
}
