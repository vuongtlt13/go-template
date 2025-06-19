package admin

import (
	"yourapp/internal/handler/admin"
	"yourapp/pkg/database"

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
	userHandler := admin.NewUserHandler(database.GetDatabase())
	api.Get("/", userHandler.GetUsers)
	api.Get("/:id", userHandler.GetUser)
	//api.Post("/", userHandler.CreateUser)
	//api.Put("/:id", userHandler.UpdateUser)
	//api.Delete("/:id", userHandler.DeleteUser)

}
