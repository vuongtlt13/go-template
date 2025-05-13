package server

import (
	"github.com/gofiber/fiber/v2"
)

// ErrorResponse represents a standardized error response
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// NewErrorResponse creates a new error response
func NewErrorResponse(code int, message string) *ErrorResponse {
	return &ErrorResponse{
		Code:    code,
		Message: message,
	}
}

// SendError sends an error response
func SendError(c *fiber.Ctx, err error) error {
	if err == nil {
		return nil
	}

	// Default to internal server error
	code := fiber.StatusInternalServerError
	message := "Internal server error"

	// Check if it's a Fiber error
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
		message = e.Message
	}

	return c.Status(code).JSON(NewErrorResponse(code, message))
}
