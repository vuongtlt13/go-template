package server

import (
	"context"
	"fmt"
	"yourapp/internal/handler"
	"yourapp/internal/handler/i18n"
	"yourapp/pb/auth/authconnect"
	i18nconnect "yourapp/pb/i18n/i18nconnect"
	"yourapp/pkg/config"
	"yourapp/pkg/server"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

// User represents the user HTTP server
type User struct {
	*server.BaseServer
}

// NewUserServer creates a new user server instance
func NewUserServer(cfg *config.Config) *User {
	return &User{
		BaseServer: server.NewBaseServer(cfg, "User"),
	}
}

// Start starts the server
func (s *User) Start() error {
	// Initialize handlers
	authHandler := handler.NewAuthHandler()

	// Create Connect service handlers
	authPath, authConnectHandler := authconnect.NewAuthServiceHandler(authHandler)

	// Create i18n Connect handler
	i18nHandler := i18n.NewHandler()
	i18nPath, i18nConnectHandler := i18nconnect.NewI18NServiceHandler(i18nHandler)

	// Mount Connect services
	app := s.GetApp()

	app.All(authPath+"*", adaptor.HTTPHandler(authConnectHandler))
	app.All(i18nPath+"*", adaptor.HTTPHandler(i18nConnectHandler))

	// Optionally, support /lang/:lang.json
	app.All("/lang/:lang.json", func(c *fiber.Ctx) error {
		c.Path("/lang/" + c.Params("lang"))
		return adaptor.HTTPHandler(i18nConnectHandler)(c)
	})

	// Start server
	return app.Listen(fmt.Sprintf(":%d", s.GetConfig().Server.Port))
}

// Shutdown gracefully shuts down the server
func (s *User) Shutdown(ctx context.Context) error {
	return s.BaseServer.Shutdown(ctx)
}
