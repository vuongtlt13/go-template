package middleware

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestRequestMiddleware(t *testing.T) {
	app := fiber.New()

	app.Use(RequestMiddleware)

	// Add a test route to check the request in context
	app.Get("/", func(c *fiber.Ctx) error {
		req := c.Locals("request")
		return c.JSON(fiber.Map{"request": req != nil})
	})

	tests := []struct {
		name           string
		method         string
		path           string
		expectedStatus int
	}{
		{
			name:           "GET request",
			method:         "GET",
			path:           "/",
			expectedStatus: fiber.StatusOK,
		},
		{
			name:           "POST request",
			method:         "POST",
			path:           "/",
			expectedStatus: fiber.StatusMethodNotAllowed,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, tt.path, nil)
			resp, err := app.Test(req)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedStatus, resp.StatusCode)

			if tt.expectedStatus == fiber.StatusOK {
				// Parse response body
				var result fiber.Map
				err = json.NewDecoder(resp.Body).Decode(&result)
				assert.NoError(t, err)

				// Check if request is stored in context
				assert.True(t, result["request"].(bool))
			}
		})
	}
}
