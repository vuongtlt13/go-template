package server

import (
	"context"
	"fmt"
	"time"

	"yourapp/internal/domain/handler"
	"yourapp/pb/auth/authconnect"
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

// UserServer represents the user HTTP server
type UserServer struct {
	app *fiber.App
	db  *gorm.DB
}

// NewUserServer creates a new user server instance
func NewUserServer() *UserServer {
	s := &UserServer{
		db: database.GetDatabase(),
	}

	// Create Fiber app
	s.app = fiber.New(fiber.Config{
		AppName:      "YourApp User",
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
func (s *UserServer) Start() error {
	// Initialize handlers
	authHandler := handler.NewAuthHandler()

	// Create Connect service handlers
	authPath, authConnectHandler := authconnect.NewAuthServiceHandler(authHandler)

	// Mount Connect services
	s.app.All(authPath+"*", adaptor.HTTPHandler(authConnectHandler))

	// Start server
	return s.app.Listen(fmt.Sprintf(":%d", config.GetConfig().Server.UserPort))
}

// Shutdown gracefully shuts down the server
func (s *UserServer) Shutdown(ctx context.Context) error {
	if s.app != nil {
		return s.app.ShutdownWithContext(ctx)
	}
	return nil
}
