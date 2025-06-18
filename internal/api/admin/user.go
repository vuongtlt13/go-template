package admin

import (
	"yourapp/internal/handler/admin"

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
	api := rootRouter.Group("/crud/user")

	// User management
	userHandler := admin.NewUserHandler()
	api.Get("/users", userHandler.GetUsers)
	api.Get("/users/:id", userHandler.GetUser)
	api.Post("/users", userHandler.CreateUser)
	api.Put("/users/:id", userHandler.UpdateUser)
	api.Delete("/users/:id", userHandler.DeleteUser)

}
