package common

import (
	"yourapp/internal/handler/common"

	"github.com/gofiber/fiber/v2"
)

type HealthRouter struct{}

func NewHealthRouter() *HealthRouter { return &HealthRouter{} }

func (r *HealthRouter) Register(rootRouter fiber.Router) {
	handler := common.NewHealthHandler()
	rootRouter.Get("/health", handler.Health)
}
