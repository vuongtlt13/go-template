package common

import (
	"github.com/gofiber/fiber/v2"
)

type ProfileHandler struct{}

func NewProfileHandler() *ProfileHandler {
	return &ProfileHandler{}
}

// GetProfile returns user profile
// @Summary Get user profile
// @Description Get the profile of the current user
// @Tags profile
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/profile [get]
func (h *ProfileHandler) GetProfile(c *fiber.Ctx) (*response.APIResponse[map[string]interface{}], error) {
	return &response.APIResponse[map[string]interface{}]{
		Success: true,
		Data:    map[string]interface{}{"message": "profile handler stub"},
		Message: "Profile fetched successfully",
	}, nil
}
