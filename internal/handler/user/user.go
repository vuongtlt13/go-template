package user

import (
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) GetProfile(c *fiber.Ctx) (interface{}, error) {
	return fiber.Map{"message": "user profile stub"}, nil
}
