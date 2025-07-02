package httperror

import (
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestNewHTTPException(t *testing.T) {
	err := NewHTTPException(fiber.StatusBadRequest, "bad request")
	if err.HttpCode != fiber.StatusBadRequest {
		t.Errorf("expected HttpCode %d, got %d", fiber.StatusBadRequest, err.HttpCode)
	}
	if err.ErrorCode != fiber.StatusBadRequest {
		t.Errorf("expected ErrorCode %d, got %d", fiber.StatusBadRequest, err.ErrorCode)
	}
	if err.Message != "bad request" {
		t.Errorf("expected Message 'bad request', got '%s'", err.Message)
	}
	if err.Stack != "" {
		t.Errorf("expected Stack '', got '%s'", err.Stack)
	}
}

func TestHTTPExceptionChain(t *testing.T) {
	err := NewBadRequest("bad req").WithErrorCode(1234).WithStack("stacktrace")
	if err.HttpCode != fiber.StatusBadRequest {
		t.Errorf("expected HttpCode %d, got %d", fiber.StatusBadRequest, err.HttpCode)
	}
	if err.ErrorCode != 1234 {
		t.Errorf("expected ErrorCode 1234, got %d", err.ErrorCode)
	}
	if err.Message != "bad req" {
		t.Errorf("expected Message 'bad req', got '%s'", err.Message)
	}
	if err.Stack != "stacktrace" {
		t.Errorf("expected Stack 'stacktrace', got '%s'", err.Stack)
	}
}

func TestCommonHTTPExceptions(t *testing.T) {
	tests := []struct {
		name     string
		fn       func(string) *HTTPException
		httpCode int
	}{
		{"BadRequest", NewBadRequest, fiber.StatusBadRequest},
		{"UnprocessableEntity", NewUnprocessableEntity, fiber.StatusUnprocessableEntity},
		{"InternalServer", NewInternalServer, fiber.StatusInternalServerError},
		{"Unauthorized", NewUnauthorized, fiber.StatusUnauthorized},
		{"Forbidden", NewForbidden, fiber.StatusForbidden},
		{"NotFound", NewNotFound, fiber.StatusNotFound},
		{"Conflict", NewConflict, fiber.StatusConflict},
		{"TooManyRequests", NewTooManyRequests, fiber.StatusTooManyRequests},
	}
	for _, tt := range tests {
		err := tt.fn("msg")
		if err.HttpCode != tt.httpCode {
			t.Errorf("%s: expected HttpCode %d, got %d", tt.name, tt.httpCode, err.HttpCode)
		}
		if err.ErrorCode != tt.httpCode {
			t.Errorf("%s: expected ErrorCode %d, got %d", tt.name, tt.httpCode, err.ErrorCode)
		}
		if err.Message != "msg" {
			t.Errorf("%s: expected Message 'msg', got '%s'", tt.name, err.Message)
		}
		if err.Stack != "" {
			t.Errorf("%s: expected Stack '', got '%s'", tt.name, err.Stack)
		}
	}
}
