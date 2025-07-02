package httperror

import (
	"github.com/gofiber/fiber/v2"
)

type HTTPException struct {
	HttpCode  int    `json:"httpCode"`
	ErrorCode int    `json:"errorCode"`
	Message   string `json:"message"`
	Stack     string `json:"stack"`
}

func NewHTTPException(httpCode int, message string) *HTTPException {
	return &HTTPException{
		HttpCode:  httpCode,
		ErrorCode: httpCode,
		Message:   message,
		Stack:     "",
	}
}

func NewBadRequest(message string) *HTTPException {
	return NewHTTPException(fiber.StatusBadRequest, message)
}

func NewUnprocessableEntity(message string) *HTTPException {
	return NewHTTPException(fiber.StatusUnprocessableEntity, message)
}

func NewInternalServer(message string) *HTTPException {
	return NewHTTPException(fiber.StatusInternalServerError, message)
}

func NewUnauthorized(message string) *HTTPException {
	return NewHTTPException(fiber.StatusUnauthorized, message)
}

func NewForbidden(message string) *HTTPException {
	return NewHTTPException(fiber.StatusForbidden, message)
}

func NewNotFound(message string) *HTTPException {
	return NewHTTPException(fiber.StatusNotFound, message)
}

func NewConflict(message string) *HTTPException {
	return NewHTTPException(fiber.StatusConflict, message)
}

func NewTooManyRequests(message string) *HTTPException {
	return NewHTTPException(fiber.StatusTooManyRequests, message)
}

func (e *HTTPException) WithErrorCode(code int) *HTTPException {
	e.ErrorCode = code
	return e
}

func (e *HTTPException) WithStack(stack string) *HTTPException {
	e.Stack = stack
	return e
}

func (e *HTTPException) Error() string {
	return e.Message
}
