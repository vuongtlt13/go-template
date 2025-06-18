package admin

import (
	"strconv"
	"yourapp/internal/model"
	"yourapp/internal/repository"
	"yourapp/internal/service"
	"yourapp/internal/validator"
	"yourapp/pkg/config"
	"yourapp/pkg/database"

	"github.com/gofiber/fiber/v2"
)

// UserHandler implements the Connect gRPC interface for admin user service
// It should satisfy adminconnect.UserServiceHandler

type UserHandler struct {
	cfg         *config.Config
	userService service.UserService
}

// NewUserHandler creates a new user handler using singleton instances
func NewUserHandler() *UserHandler {
	db := database.GetDatabase()
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(db, userRepo)
	cfg := config.GetConfig()
	return &UserHandler{
		cfg:         cfg,
		userService: userService,
	}
}

// GetUsers returns list of users (admin view)
// @Summary Get users (admin)
// @Description Get list of users with admin privileges
// @Tags admin
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Items per page"
// @Success 200 {object} AdminUsersResponse
// @Failure 401 {object} ErrorResponse
// @Router /api/v1/admin/users [get]
func (h *UserHandler) GetUsers(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))

	users, total, err := h.userService.ListUsers(c.Context(), page, limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get users",
		})
	}

	totalPages := (total + int64(limit) - 1) / int64(limit)

	return c.JSON(AdminUsersResponse{
		Users:      users,
		Total:      total,
		Page:       page,
		Limit:      limit,
		TotalPages: int(totalPages),
	})
}

// GetUser returns a single user (admin view)
// @Summary Get user (admin)
// @Description Get user details with admin privileges
// @Tags admin
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} AdminUserResponse
// @Failure 401 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/v1/admin/users/{id} [get]
func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	userID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}
	user, err := h.userService.GetUserByID(c.Context(), uint(userID))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	return c.JSON(user)
}

// CreateUser creates a new user (admin)
// @Summary Create user (admin)
// @Description Create a new user with admin privileges
// @Tags admin
// @Accept json
// @Produce json
// @Param user body CreateUserRequest true "User data"
// @Success 201 {object} AdminUserResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Router /api/v1/admin/users [post]
func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req CreateUserRequest
	if err := validator.ValidateRequest(c, &req); err != nil {
		return err
	}

	user, err := h.userService.CreateUser(c.Context(), &model.User{
		Email:     req.Email,
		Password:  req.Password,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Role:      req.Role,
	})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

// UpdateUser updates user information (admin)
// @Summary Update user (admin)
// @Description Update user information with admin privileges
// @Tags admin
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body UpdateUserRequest true "User data"
// @Success 200 {object} AdminUserResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/v1/admin/users/{id} [put]
func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var req UpdateUserRequest
	if err := validator.ValidateRequest(c, &req); err != nil {
		return err
	}

	user, err := h.userService.UpdateUser(c.Context(), id, &req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(user)
}

// DeleteUser deletes a user (admin)
// @Summary Delete user (admin)
// @Description Delete user with admin privileges
// @Tags admin
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/v1/admin/users/{id} [delete]
func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.userService.DeleteUser(c.Context(), id); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "User deleted successfully",
	})
}
