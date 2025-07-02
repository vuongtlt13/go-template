package response

import "github.com/gofiber/fiber/v2"

func SuccessResponse(c *fiber.Ctx, data interface{}, message string) error {
	return c.Status(fiber.StatusOK).JSON(APIResponse[interface{}]{
		Success: true,
		Data:    data,
		Message: message,
	})
}
