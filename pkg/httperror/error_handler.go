package httperror

import (
	"errors"
	"yourapp/pkg/logger"
	"yourapp/pkg/response"

	"github.com/gofiber/fiber/v2"
)

var _logger = logger.GetLogger()

func HandlerError(c *fiber.Ctx, err error) error {
	var httpErr *HTTPException
	if errors.As(err, &httpErr) {
		return response.ErrorResponse(
			c,
			httpErr.HttpCode,
			httpErr.Message,
			httpErr.ErrorCode,
			httpErr.Stack,
		)
	}
	var fiberError *fiber.Error
	if errors.As(err, &fiberError) {
		return response.ErrorResponse(
			c,
			fiberError.Code,
			fiberError.Message,
			fiberError.Code,
		)
	}
	_logger.Error("error %v", err)
	return response.ErrorResponse(
		c,
		fiber.StatusInternalServerError,
		"unknown_error",
		999,
		err.Error(),
	)
}
