package response

import (
	"github.com/gofiber/fiber/v2"
)

// To use HTTP errors, import from "pkg/httperror"

func ErrorResponse(c *fiber.Ctx, status int, message string, errorCode int, stack ...string) error {
	var stackVal string
	if len(stack) > 0 {
		stackVal = stack[0]
	}
	return c.Status(status).JSON(APIResponse[interface{}]{
		Success:   false,
		Data:      nil,
		Message:   message,
		ErrorCode: errorCode,
		Stack:     stackVal,
	})
}
