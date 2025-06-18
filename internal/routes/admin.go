package routes

import (
	"yourapp/internal/api/admin"

	"github.com/gofiber/fiber/v2"
)

// Router represents admin routes
type AdminRouter struct {
	app *fiber.App
}

// NewRouter creates a new admin router
func NewAdminRouter() *AdminRouter {
	return &AdminRouter{
		app: nil,
	}
}

// Register registers all admin routes
func (r *AdminRouter) Register(app *fiber.App) {
	r.app = app
	api := r.app.Group("/api")

	userRouter := admin.NewUserRouter()
	userRouter.Register(api)

	// // User management
	// userHandler := NewUserHandler(r.cfg, r.adminService)
	// admin.Get("/users", userHandler.GetUsers)
	// admin.Get("/users/:id", userHandler.GetUser)
	// admin.Post("/users", userHandler.CreateUser)
	// admin.Put("/users/:id", userHandler.UpdateUser)
	// admin.Delete("/users/:id", userHandler.DeleteUser)

}
