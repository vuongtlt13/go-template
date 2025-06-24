package common

import (
	"yourapp/internal/handler/common"
	"yourapp/internal/service"
	"yourapp/pkg/config"

	"github.com/gofiber/fiber/v2"
)

type AuthRouter struct {
	cfg         *config.Config
	authService service.AuthService
}

func NewAuthRouter(cfg *config.Config, authService service.AuthService) *AuthRouter {
	return &AuthRouter{cfg: cfg, authService: authService}
}

func (r *AuthRouter) Register(rootRouter fiber.Router) {
	handler := common.NewAuthHandler(r.cfg, r.authService)
	rootRouter.Post("/auth/login", handler.Login)
}
