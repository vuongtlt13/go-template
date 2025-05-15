package server

import (
	"context"
	"fmt"
	"yourapp/internal/handler"
	adminhandler "yourapp/internal/handler/admin"
	"yourapp/internal/handler/i18n"
	"yourapp/internal/repository"
	"yourapp/internal/service"
	"yourapp/pb/admin/adminconnect"
	"yourapp/pb/auth/authconnect"
	i18nconnect "yourapp/pb/i18n/i18nconnect"
	"yourapp/pkg/config"
	"yourapp/pkg/database"
	"yourapp/pkg/server"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"gorm.io/gorm"
)

// Admin represents the admin HTTP server
type Admin struct {
	*server.BaseServer
	db *gorm.DB
}

// NewAdminServer creates a new admin server instance
func NewAdminServer(cfg *config.Config) *Admin {
	return &Admin{
		BaseServer: server.NewBaseServer(cfg, "Admin"),
		db:         database.GetDatabase(),
	}
}

// Start starts the server
func (s *Admin) Start() error {
	// Initialize handlers
	authHandler := handler.NewAuthHandler()
	userHandler := adminhandler.NewUserHandler(s.db)

	// Initialize repositories
	roleRepo := repository.NewRoleRepository()
	permissionRepo := repository.NewPermissionRepository()

	// Initialize services
	roleService := service.NewRoleService(s.db, roleRepo, permissionRepo)

	// Initialize handlers with services
	roleHandler := adminhandler.NewRoleHandler(roleService)
	permissionHandler := adminhandler.NewPermissionHandler()

	// Create Connect service handlers
	authPath, authConnectHandler := authconnect.NewAuthServiceHandler(authHandler)
	userPath, userConnectHandler := adminconnect.NewUserServiceHandler(userHandler)
	rolePath, roleConnectHandler := adminconnect.NewRoleServiceHandler(roleHandler)
	permissionPath, permissionConnectHandler := adminconnect.NewPermissionServiceHandler(permissionHandler)

	// Create i18n Connect handler
	i18nHandler := i18n.NewHandler()
	i18nPath, i18nConnectHandler := i18nconnect.NewI18NServiceHandler(i18nHandler)

	// Mount Connect services
	app := s.GetApp()

	app.All(authPath+"*", adaptor.HTTPHandler(authConnectHandler))
	app.All(userPath+"*", adaptor.HTTPHandler(userConnectHandler))
	app.All(rolePath+"*", adaptor.HTTPHandler(roleConnectHandler))
	app.All(permissionPath+"*", adaptor.HTTPHandler(permissionConnectHandler))
	app.All(i18nPath+"*", adaptor.HTTPHandler(i18nConnectHandler))

	// Optionally, support /lang/:lang.json
	app.All("/lang/:lang.json", func(c *fiber.Ctx) error {
		c.Path("/lang/" + c.Params("lang"))
		return adaptor.HTTPHandler(i18nConnectHandler)(c)
	})

	// Start server
	return app.Listen(fmt.Sprintf(":%d", s.GetConfig().Server.AdminPort))
}

// Shutdown gracefully shuts down the server
func (s *Admin) Shutdown(ctx context.Context) error {
	return s.BaseServer.Shutdown(ctx)
}
