package common

import (
	"yourapp/internal/handler/common"

	autofiber "github.com/vuongtlt13/auto-fiber"
)

type HealthRouter struct{}

func NewHealthRouter() *HealthRouter { return &HealthRouter{} }

func (r *HealthRouter) Register(rootRouter *autofiber.AutoFiberGroup) {
	handler := common.NewHealthHandler()
	rootRouter.Get("/health", handler.Health,
		autofiber.WithDescription("Get server health status"),
		autofiber.WithTags("health"),
	)
}
