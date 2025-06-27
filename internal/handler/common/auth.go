package common

import (
	"yourapp/internal/schema"
	"yourapp/internal/service"
	"yourapp/pkg/config"
	"yourapp/pkg/response"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	cfg     *config.Config
	service service.AuthService
}

func NewAuthHandler(cfg *config.Config, service service.AuthService) *AuthHandler {
	return &AuthHandler{
		cfg:     cfg,
		service: service,
	}
}

// Login handles user login
// @Summary Login user
// @Description Authenticate user and return JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body schema.LoginRequest true "Login credentials"
// @Success 200 {object} schema.APIResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Router /api/v1/auth/login [post]
func (h *AuthHandler) Login(c *fiber.Ctx, req schema.LoginRequest) error {
	token, err := h.service.Login(c.Context(), service.Credential{
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		return response.ErrorResponse(c, fiber.StatusUnauthorized, err.Error(), fiber.StatusUnauthorized)
	}

	return response.SuccessResponse(c, schema.LoginData{Token: token}, "Login successful")
}

// Register handles user registration
// @Summary Register new user
// @Description Create a new user account
// @Tags auth
// @Accept json
// @Produce json
// @Param user body schema.RegisterRequest true "User registration data"
// @Success 201 {object} schema.APIResponse
// @Failure 400 {object} ErrorResponse
// @Router /api/v1/auth/register [post]
func (h *AuthHandler) Register(c *fiber.Ctx, req schema.RegisterRequest) error {
	err := h.service.Register(c.Context(), service.Credential{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return response.ErrorResponse(c, fiber.StatusBadRequest, err.Error(), fiber.StatusBadRequest)
	}

	return response.SuccessResponse(c, schema.RegisterData{Message: "User registered successfully"}, "Registration successful")
}
