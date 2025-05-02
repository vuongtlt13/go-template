package http

import (
	"github.com/gofiber/fiber/v2"
	"yourapp/internal/service"
	"yourapp/pkg/auth"
	"yourapp/pkg/logger"
)

type Handler struct {
	userService service.UserService
	authService service.AuthService
	jwtManager  *auth.JWTManager
	logger      logger.Logger
}

func NewHandler(userService service.UserService, authService service.AuthService, jwtManager *auth.JWTManager) *Handler {
	return &Handler{
		userService: userService,
		authService: authService,
		jwtManager:  jwtManager,
		logger:      logger.GetLogger(),
	}
}

func (h *Handler) SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v1")

	// Public routes
	authRouter := api.Group("/authRouter")
	authRouter.Post("/register", h.Register)
	authRouter.Post("/login", h.Login)
	authRouter.Get("/verify", h.VerifyEmail)

	// Protected routes
	users := api.Group("/users")
	users.Use(h.AuthMiddleware)
	users.Get("/", h.ListUsers)
	users.Get("/:id", h.GetUser)
	users.Put("/:id", h.UpdateUser)
	users.Delete("/:id", h.DeleteUser)
}

func (h *Handler) AuthMiddleware(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		return fiber.NewError(fiber.StatusUnauthorized, "missing authorization header")
	}

	email, err := h.jwtManager.VerifyToken(token)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "invalid token")
	}

	c.Locals("email", email)
	return c.Next()
}

func (h *Handler) Register(c *fiber.Ctx) error {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		FullName string `json:"full_name"`
	}

	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}

	user, err := h.userService.Register(c.Context(), req.Email, req.Password, req.FullName)
	if err != nil {
		h.logger.Error("Failed to register user", "error", err)
		return err
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func (h *Handler) Login(c *fiber.Ctx) error {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}

	token, user, err := h.userService.Login(c.Context(), req.Email, req.Password)
	if err != nil {
		h.logger.Error("Failed to login user", "error", err)
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(map[string]interface{}{
		"token": token,
		"user":  user,
	})
}

func (h *Handler) VerifyEmail(c *fiber.Ctx) error {
	token := c.Query("token")
	if token == "" {
		return fiber.NewError(fiber.StatusInternalServerError, "missing token")
	}

	if err := h.userService.VerifyEmail(c.Context(), token); err != nil {
		h.logger.Error("Failed to verify email", "error", err)
		return err
	}

	return c.SendString("Email verified successfully")
}

func (h *Handler) ListUsers(c *fiber.Ctx) error {
	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 10)

	users, total, err := h.userService.ListUsers(c.Context(), page, limit)
	if err != nil {
		h.logger.Error("Failed to list users", "error", err)
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(map[string]interface{}{
		"users": users,
		"total": total,
	})
}

func (h *Handler) GetUser(c *fiber.Ctx) error {
	id := c.Params("id")

	user, err := h.userService.GetUser(c.Context(), id)
	if err != nil {
		h.logger.Error("Failed to get user", "error", err)
		return err
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func (h *Handler) UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")

	var req struct {
		FullName string `json:"full_name"`
		Role     string `json:"role"`
	}

	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}

	user, err := h.userService.UpdateUser(c.Context(), id, req.FullName, req.Role)
	if err != nil {
		h.logger.Error("Failed to update user", "error", err)
		return err
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func (h *Handler) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := h.userService.DeleteUser(c.Context(), id); err != nil {
		h.logger.Error("Failed to delete user", "error", err)
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}
