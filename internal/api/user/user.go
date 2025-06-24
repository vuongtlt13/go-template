package user

import (
	"yourapp/internal/handler/user"

	"github.com/gofiber/fiber/v2"
)

// Router represents user routes
type UserRouter struct {
	app *fiber.App
}

// NewUserRouter creates a new user router
func NewUserRouter() *UserRouter {
	return &UserRouter{
		app: nil,
	}
}

// Register registers all user routes
func (r *UserRouter) Register(rootRouter fiber.Router) {
	api := rootRouter.Group("/user")

	// User management
	userHandler := user.NewUserHandler()
	api.Get("/profile", userHandler.GetProfile)
	// Thêm các route user khác nếu cần
}
