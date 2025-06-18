package common

import (
	"yourapp/internal/service"
	"yourapp/internal/validator"
	"yourapp/pkg/config"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	cfg     *config.Config
	service *service.AuthService
}

func NewAuthHandler(cfg *config.Config, service *service.AuthService) *AuthHandler {
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
// @Param credentials body LoginRequest true "Login credentials"
// @Success 200 {object} LoginResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Router /api/v1/auth/login [post]
func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req LoginRequest
	if err := validator.ValidateRequest(c, &req); err != nil {
		return err
	}

	token, err := h.service.Login(c.Context(), &req)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(token)
}

// Register handles user registration
// @Summary Register new user
// @Description Create a new user account
// @Tags auth
// @Accept json
// @Produce json
// @Param user body RegisterRequest true "User registration data"
// @Success 201 {object} RegisterResponse
// @Failure 400 {object} ErrorResponse
// @Router /api/v1/auth/register [post]
func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var req RegisterRequest
	if err := validator.ValidateRequest(c, &req); err != nil {
		return err
	}

	user, err := h.service.Register(c.Context(), &req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

// RefreshToken handles token refresh
// @Summary Refresh JWT token
// @Description Get a new access token using refresh token
// @Tags auth
// @Accept json
// @Produce json
// @Param refresh_token body RefreshTokenRequest true "Refresh token"
// @Success 200 {object} RefreshTokenResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Router /api/v1/auth/refresh [post]
func (h *AuthHandler) RefreshToken(c *fiber.Ctx) error {
	var req RefreshTokenRequest
	if err := validator.ValidateRequest(c, &req); err != nil {
		return err
	}

	token, err := h.service.RefreshToken(c.Context(), req.RefreshToken)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(token)
}

// Logout handles user logout
// @Summary Logout user
// @Description Invalidate user's JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {object} SuccessResponse
// @Failure 401 {object} ErrorResponse
// @Router /api/v1/auth/logout [post]
func (h *AuthHandler) Logout(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if err := h.service.Logout(c.Context(), token); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Successfully logged out",
	})
}
