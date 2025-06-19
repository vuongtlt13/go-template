package response

import (
	"errors"
	"yourapp/pkg/logger"

	"github.com/gofiber/fiber/v2"
)

var _logger = logger.GetLogger()

func HandlerError(c *fiber.Ctx, err error) error {
	var fiberError *fiber.Error
	if errors.As(err, &fiberError) {
		return ErrorResponse(c, err.(*fiber.Error).Code, err.(*fiber.Error).Message, err.(*fiber.Error).Code)
	}
	_logger.Error("error %v", err)
	return ErrorResponse(c, fiber.StatusInternalServerError, err.Error(), 9999)
}
