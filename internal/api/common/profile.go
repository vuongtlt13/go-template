package common

import (
	"yourapp/pkg/response"

	"github.com/gofiber/fiber/v2"
)

type ProfileRouter struct{}

func NewProfileRouter() *ProfileRouter { return &ProfileRouter{} }

func (r *ProfileRouter) Register(rootRouter fiber.Router) {
	rootRouter.Get("/profile", func(c *fiber.Ctx) error {
		return response.SuccessResponse(c, fiber.Map{"message": "profile stub"}, "profile stub")
	})
}
