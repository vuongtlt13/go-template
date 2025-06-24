package common

import (
	"yourapp/internal/handler/common"

	"github.com/gofiber/fiber/v2"
)

type I18nRouter struct{}

func NewI18nRouter() *I18nRouter { return &I18nRouter{} }

func (r *I18nRouter) Register(rootRouter fiber.Router) {
	handler := common.NewI18nHandler()
	rootRouter.Get("/i18n", handler.GetI18n)
}
