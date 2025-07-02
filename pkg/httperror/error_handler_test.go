package httperror

import (
	"encoding/json"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

type mockLogger struct{}

func (m *mockLogger) Error(format string, args ...interface{}) {}

func TestHandlerError_HTTPException(t *testing.T) {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return HandlerError(c, NewBadRequest("bad req").WithErrorCode(1234).WithStack("stacktrace"))
	})
	resp, _ := app.Test(httptest.NewRequest("GET", "/", nil))
	var body map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&body)
	assert.Equal(t, float64(400), body["errorCode"])
	assert.Equal(t, "bad req", body["message"])
	assert.Equal(t, "stacktrace", body["stack"])
	assert.Equal(t, false, body["success"])
}

func TestHandlerError_FiberError(t *testing.T) {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return HandlerError(c, fiber.NewError(fiber.StatusUnauthorized, "unauthorized"))
	})
	resp, _ := app.Test(httptest.NewRequest("GET", "/", nil))
	var body map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&body)
	assert.Equal(t, float64(401), body["errorCode"])
	assert.Equal(t, "unauthorized", body["message"])
	assert.Equal(t, false, body["success"])
}

func TestHandlerError_UnknownError(t *testing.T) {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return HandlerError(c, errors.New("something went wrong"))
	})
	resp, _ := app.Test(httptest.NewRequest("GET", "/", nil))
	var body map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&body)
	assert.Equal(t, float64(999), body["errorCode"])
	assert.Equal(t, "unknown_error", body["message"])
	assert.Equal(t, false, body["success"])
	assert.Equal(t, "something went wrong", body["stack"])
}
