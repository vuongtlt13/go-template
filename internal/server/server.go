package server

import (
	"fmt"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"time"
	"yourapp/internal/domain/handler"
	"yourapp/pb/auth/authconnect"
	"yourapp/pb/health/healthconnect"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"yourapp/pkg/config"
)

type Server struct {
	fiber *fiber.App
	cfg   *config.Config
}

func NewServer() *Server {
	// Load configuration
	cfg := config.GetConfig()

	// Setup HTTP server
	app := fiber.New(fiber.Config{
		ErrorHandler: errorHandler,
	})

	// Middleware
	// 🛡️ Recover from panics (should be first)
	app.Use(recover.New())

	// 🧢 Secure HTTP headers
	if cfg.IsProduction() {
		app.Use(helmet.New())
	}

	// 🌐 Allow cross-origin requests (adjust config for production)
	allowOrigins := "*"
	if cfg.IsProduction() {
		allowOrigins = cfg.Cors
	}
	app.Use(cors.New(cors.Config{
		AllowOrigins: allowOrigins,
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
	}))

	// 🚦 Rate limit (e.g., 20 requests per minute per IP)
	app.Use(limiter.New(limiter.Config{
		Max:        60,
		Expiration: 1 * time.Minute,
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"message": "Too many requests. Please try again later.",
			})
		},
	}))

	// Initialize your Connect service implementation
	authHandler := handler.NewAuthHandler()
	// Create the Connect HTTP handler
	path, connectHandler := authconnect.NewAuthServiceHandler(authHandler)
	// Mount Connect service using adaptor
	app.All(path+"*", adaptor.HTTPHandler(connectHandler))

	// Initialize your Connect service implementation
	healthHandler := handler.NewHealthHandler()
	// Create the Connect HTTP handler
	path, connectHandler = healthconnect.NewHealthServiceHandler(healthHandler)
	// Mount Connect service using adaptor
	app.All(path+"*", adaptor.HTTPHandler(connectHandler))

	return &Server{
		fiber: app,
		cfg:   cfg,
	}
}

func (s *Server) Start() error {
	err := s.fiber.Listen(fmt.Sprintf(":%d", s.cfg.Server.Port))
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) Shutdown() error {
	err := s.fiber.Shutdown()
	if err != nil {
		return err
	}
	return nil
}
