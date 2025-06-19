package common

import (
	"yourapp/internal/service"
	"yourapp/pkg/config"

	"github.com/gofiber/fiber/v2"
)

// Router represents common routes
type Router struct {
	app         *fiber.App
	cfg         *config.Config
	authService service.AuthService
}

// NewRouter creates a new common router
func NewRouter(app *fiber.App, cfg *config.Config, authService service.AuthService) *Router {
	return &Router{
		app:         app,
		cfg:         cfg,
		authService: authService,
	}
}

// Register registers all common routes
func (r *Router) Register() {
	api := r.app.Group("/api/v1")

	// Health check
	healthHandler := NewHealthHandler()
	api.Get("/health", healthHandler.HealthCheck)

	// Auth routes
	authHandler := NewAuthHandler(r.cfg, r.authService)
	auth := api.Group("/auth")
	auth.Post("/login", authHandler.Login)
	auth.Post("/register", authHandler.Register)
}
