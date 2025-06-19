package response

import "github.com/gofiber/fiber/v2"

type successResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}

func SuccessResponse(c *fiber.Ctx, data interface{}, message string) error {
	return c.Status(fiber.StatusOK).JSON(successResponse{
		Success: true,
		Data:    data,
		Message: message,
	})
}
