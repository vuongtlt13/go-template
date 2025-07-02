package common

import (
	"yourapp/internal/handler/common"
	"yourapp/internal/schema"
	"yourapp/internal/service"
	"yourapp/pkg/config"

	autofiber "github.com/vuongtlt13/auto-fiber"
)

type AuthRouter struct {
	cfg         *config.Config
	authService service.AuthService
}

func NewAuthRouter(cfg *config.Config, authService service.AuthService) *AuthRouter {
	return &AuthRouter{cfg: cfg, authService: authService}
}

func (r *AuthRouter) Register(rootRouter *autofiber.AutoFiberGroup) {
	authHandler := common.NewAuthHandler(r.cfg, r.authService)

	rootRouter.Post("/auth/login", authHandler.Login,
		autofiber.WithDescription("Authenticate user and return JWT token"),
		autofiber.WithTags("auth"),
		autofiber.WithRequestSchema(schema.LoginRequest{}),
		autofiber.WithResponseSchema(response.APIResponse[schema.LoginData]{}),
	)

	rootRouter.Post("/auth/register", authHandler.Register,
		autofiber.WithDescription("Register new user account"),
		autofiber.WithTags("auth"),
		autofiber.WithRequestSchema(schema.RegisterRequest{}),
		autofiber.WithResponseSchema(response.APIResponse[schema.RegisterData]{}),
	)
}
