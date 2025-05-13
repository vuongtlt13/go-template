package server

import (
	"context"
	"fmt"
	"yourapp/internal/domain/handler"
	adminhandler "yourapp/internal/domain/handler/admin"
	"yourapp/pb/admin/adminconnect"
	"yourapp/pb/auth/authconnect"
	"yourapp/pkg/config"
	"yourapp/pkg/database"
	"yourapp/pkg/server"

	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"gorm.io/gorm"
)

// AdminServer represents the admin HTTP server
type AdminServer struct {
	*server.BaseServer
	db *gorm.DB
}

// NewAdminServer creates a new admin server instance
func NewAdminServer(cfg *config.Config) *AdminServer {
	return &AdminServer{
		BaseServer: server.NewBaseServer(cfg, "Admin"),
		db:         database.GetDatabase(),
	}
}

// Start starts the server
func (s *AdminServer) Start() error {
	// Initialize handlers
	authHandler := handler.NewAuthHandler()
	userHandler := adminhandler.NewUserHandler(s.db)
	roleHandler := adminhandler.NewRoleHandler(s.db)
	permissionHandler := adminhandler.NewPermissionHandler()

	// Create Connect service handlers
	authPath, authConnectHandler := authconnect.NewAuthServiceHandler(authHandler)
	userPath, userConnectHandler := adminconnect.NewUserServiceHandler(userHandler)
	rolePath, roleConnectHandler := adminconnect.NewRoleServiceHandler(roleHandler)
	permissionPath, permissionConnectHandler := adminconnect.NewPermissionServiceHandler(permissionHandler)

	// Mount Connect services
	app := s.GetApp()
	app.All(authPath+"*", adaptor.HTTPHandler(authConnectHandler))
	app.All(userPath+"*", adaptor.HTTPHandler(userConnectHandler))
	app.All(rolePath+"*", adaptor.HTTPHandler(roleConnectHandler))
	app.All(permissionPath+"*", adaptor.HTTPHandler(permissionConnectHandler))

	// Start server
	return app.Listen(fmt.Sprintf(":%d", s.GetConfig().Server.AdminPort))
}

// Shutdown gracefully shuts down the server
func (s *AdminServer) Shutdown(ctx context.Context) error {
	return s.BaseServer.Shutdown(ctx)
}

// Additional admin-specific methods
