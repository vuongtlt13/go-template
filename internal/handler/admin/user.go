package admin

import (
	"yourapp/internal/datatable"
	internal_schema "yourapp/internal/schema"
	"yourapp/pkg/response"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// UserHandler handles user-related requests
type UserHandler struct {
	db *gorm.DB
}

// NewUserHandler creates a new user handler
func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{db: db}
}

// GetUsers handles GET /crud/user/ - returns user list with datatable support
// @Summary Get user list (datatable)
// @Description Retrieve a list of users with datatable support (pagination, filtering, etc.)
// @Tags admin-user
// @Accept json
// @Produce json
// @Success 200 {object} schema.UserDatatableResponse
// @Failure 400 {object} common.ErrorResponse
// @Router /api/crud/user/ [get]
func (h *UserHandler) GetUsers(c *fiber.Ctx) error {
	// Use the datatable helper to process the request
	userDatatable := datatable.NewUserDataTable(nil)
	return userDatatable.Render(c, nil)
	//res := []model.User{}
	//return response.SuccessResponse(c, res, "ok")
}

// GetUser handles GET /crud/user/:id - returns a specific user
// @Summary Get user details
// @Description Retrieve detailed information of a user by ID
// @Tags admin-user
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} schema.UserResponse
// @Failure 404 {object} common.ErrorResponse
// @Router /api/crud/user/{id} [get]
func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	id := c.Params("id")

	var user internal_schema.UserInfo
	if err := h.db.First(&user, id).Error; err != nil {
		return response.ErrorResponse(c, fiber.StatusNotFound, "User not found", fiber.StatusNotFound)
	}

	return response.SuccessResponse(c, &user, "ok")
}
