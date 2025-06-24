package common_test

import (
	"context"
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"yourapp/internal/handler/common"
	"yourapp/internal/service"
	"yourapp/pkg/config"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

// Mock service.AuthService
type mockAuthService struct{}

func (m *mockAuthService) Login(ctx context.Context, cred service.Credential) (string, error) {
	return "mock-token", nil
}
func (m *mockAuthService) Register(ctx context.Context, cred service.Credential) error {
	return nil
}

func TestAuthHandler_Login(t *testing.T) {
	app := fiber.New()
	cfg := &config.Config{}
	mockService := &mockAuthService{}
	handler := common.NewAuthHandler(cfg, mockService)
	app.Post("/auth/login", handler.Login)

	body := `{"email":"test@example.com","password":"123456"}`
	req := httptest.NewRequest("POST", "/auth/login", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	respBody, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)
	assert.Contains(t, string(respBody), "mock-token")
}
