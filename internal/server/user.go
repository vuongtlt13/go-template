package server

import (
	"context"
	"fmt"

	"github.com/gofiber/swagger"

	"yourapp/internal/handler/common"
	"yourapp/internal/repository"
	"yourapp/internal/service"
	"yourapp/pkg/auth"
	"yourapp/pkg/config"
	"yourapp/pkg/logger"
	"yourapp/pkg/server"
)

// User represents the user HTTP server
type User struct {
	*server.BaseServer
	logger      logger.Logger
	authService service.AuthService
}

// NewUserServer creates a new user server instance
func NewUserServer(cfg *config.Config, logger logger.Logger) *User {
	// Create repositories
	userRepo := repository.NewUserRepository(nil) // TODO: Pass DB connection

	// Create JWT manager
	jwtManager := auth.NewJWTManager(cfg.JWT.Secret, cfg.JWT.ExpirePeriod)

	// Create services
	authService := service.NewAuthService(nil, userRepo, jwtManager) // TODO: Pass DB connection

	return &User{
		BaseServer:  server.NewBaseServer(cfg, "User"),
		logger:      logger,
		authService: authService,
	}
}

// Start starts the server
func (s *User) Start() error {
	app := s.GetApp()

	// Initialize and register common routes
	commonRouter := common.NewRouter(app, s.GetConfig(), s.authService)
	commonRouter.Register()

	// Swagger documentation
	app.Get("/swagger/*", swagger.HandlerDefault)

	// Start server
	addr := fmt.Sprintf(":%d", s.GetConfig().Server.Port)
	s.logger.Info("Starting user server", "addr", addr)
	return app.Listen(addr)
}

// Shutdown gracefully shuts down the server
func (s *User) Shutdown(ctx context.Context) error {
	return s.GetApp().ShutdownWithContext(ctx)
}
