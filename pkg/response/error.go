package response

import (
	"github.com/gofiber/fiber/v2"
)

type errorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func ErrorResponse(c *fiber.Ctx, status int, message string, errorCode int) error {
	return c.Status(status).JSON(errorResponse{
		Success: false,
		Message: message,
		Code:    errorCode,
	})
}
