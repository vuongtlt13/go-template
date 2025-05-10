package server

import (
	"context"
	"fmt"
	"time"
	"yourapp/pb/admin/adminconnect"
	"yourapp/pb/auth/authconnect"

	"yourapp/internal/domain/handler"
	adminhandler "yourapp/internal/domain/handler/admin"
	"yourapp/pkg/config"
	"yourapp/pkg/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"gorm.io/gorm"
)

// AdminServer represents the admin HTTP server
type AdminServer struct {
	app *fiber.App
	db  *gorm.DB
}

// NewAdminServer creates a new admin server instance
func NewAdminServer() *AdminServer {
	s := &AdminServer{
		db: database.GetDatabase(),
	}

	// Create Fiber app
	s.app = fiber.New(fiber.Config{
		AppName:      "YourApp Admin",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	})

	// Add middleware
	s.app.Use(recover.New())
	s.app.Use(logger.New())

	// Configure CORS
	cfg := config.GetConfig()
	allowOrigins := "*"
	if cfg.IsProduction() {
		allowOrigins = cfg.Cors
	}
	s.app.Use(cors.New(cors.Config{
		AllowOrigins: allowOrigins,
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	// Add rate limiting
	s.app.Use(limiter.New(limiter.Config{
		Max:        60,
		Expiration: 1 * time.Minute,
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"message": "Too many requests. Please try again later.",
			})
		},
	}))

	return s
}

// Start starts the server
func (s *AdminServer) Start() error {
	// Initialize handlers
	authHandler := handler.NewAuthHandler()
	userHandler := adminhandler.NewUserHandler()
	roleHandler := adminhandler.NewRoleHandler()
	permissionHandler := adminhandler.NewPermissionHandler()

	// Create Connect service handlers
	authPath, authConnectHandler := authconnect.NewAuthServiceHandler(authHandler)
	userPath, userConnectHandler := adminconnect.NewUserServiceHandler(userHandler)
	rolePath, roleConnectHandler := adminconnect.NewRoleServiceHandler(roleHandler)
	permissionPath, permissionConnectHandler := adminconnect.NewPermissionServiceHandler(permissionHandler)

	// Mount Connect services
	s.app.All(authPath+"*", adaptor.HTTPHandler(authConnectHandler))
	s.app.All(userPath+"*", adaptor.HTTPHandler(userConnectHandler))
	s.app.All(rolePath+"*", adaptor.HTTPHandler(roleConnectHandler))
	s.app.All(permissionPath+"*", adaptor.HTTPHandler(permissionConnectHandler))

	// Start server
	return s.app.Listen(fmt.Sprintf(":%d", config.GetConfig().Server.AdminPort))
}

// Shutdown gracefully shuts down the server
func (s *AdminServer) Shutdown(ctx context.Context) error {
	if s.app != nil {
		return s.app.ShutdownWithContext(ctx)
	}
	return nil
}
