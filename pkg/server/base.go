package server

import (
	"context"
	"log"
	"yourapp/pkg/config"
	"yourapp/pkg/middleware"
	"yourapp/pkg/response"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	autofiber "github.com/vuongtlt13/auto-fiber"
)

// BaseServer defines common functionality for all servers
type BaseServer struct {
	app *autofiber.AutoFiber
	cfg *config.Config
}

// NewBaseServer creates a new base server with the given configuration
func NewBaseServer(cfg *config.Config, serverType string) *BaseServer {
	s := &BaseServer{
		cfg: cfg,
	}

	s.app = autofiber.New(fiber.Config{
		AppName:      cfg.Server.App.Name + " " + serverType,
		ReadTimeout:  cfg.Server.App.ReadTimeout,
		WriteTimeout: cfg.Server.App.WriteTimeout,
		IdleTimeout:  cfg.Server.App.IdleTimeout,
		ErrorHandler: response.HandlerError,
	})

	// Add middleware
	s.app.Use(recover.New())
	s.app.Use(logger.New())
	s.app.Use(middleware.RequestMiddleware) // Store request in context
	s.app.Use(middleware.I18nMiddleware)    // Handle internationalization

	// Configure CORS
	allowOrigins := "*"
	if cfg.IsProduction() {
		allowOrigins = cfg.Server.Cors
	}
	s.app.Use(cors.New(cors.Config{
		AllowOrigins: allowOrigins,
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	// Add rate limiting
	s.app.Use(limiter.New(limiter.Config{
		Max:        cfg.Server.App.RateLimit.Max,
		Expiration: cfg.Server.App.RateLimit.Expiration,
		LimitReached: func(c *fiber.Ctx) error {
			return response.ErrorResponse(c, fiber.StatusTooManyRequests, "Too many requests. Please try again later.", fiber.StatusTooManyRequests)
		},
	}))

	return s
}

// GetApp returns the AutoFiber app instance
func (s *BaseServer) GetApp() *autofiber.AutoFiber {
	return s.app
}

// GetConfig returns the server configuration
func (s *BaseServer) GetConfig() *config.Config {
	return s.cfg
}

// Shutdown gracefully shuts down the server
func (s *BaseServer) Shutdown(ctx context.Context) error {
	log.Println("Server shutting down...")
	return s.app.App.ShutdownWithContext(ctx)
}
