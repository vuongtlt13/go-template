package admin

import (
	"yourapp/internal/handler/admin"
	"yourapp/internal/schema"
	"yourapp/pkg/database"

	autofiber "github.com/vuongtlt13/auto-fiber"
)

// Router represents user routes
type UserRouter struct {
	app *autofiber.AutoFiber
}

// NewUserRouter creates a new user router
func NewUserRouter() *UserRouter {
	return &UserRouter{
		app: nil,
	}
}

// Register registers all user routes
func (r *UserRouter) Register(rootRouter *autofiber.AutoFiberGroup) {
	// User management
	userHandler := admin.NewUserHandler(database.GetDatabase())

	rootRouter.Get("/crud/user", userHandler.GetUsers,
		autofiber.WithDescription("Get all users with pagination"),
		autofiber.WithTags("admin", "user"),
		autofiber.WithRequestSchema(schema.UserQuery{}),
		autofiber.WithResponseSchema(schema.APIResponse{}),
	)

	rootRouter.Get("/crud/user/:id", userHandler.GetUser,
		autofiber.WithDescription("Get user by ID"),
		autofiber.WithTags("admin", "user"),
		autofiber.WithRequestSchema(schema.UserIDParam{}),
		autofiber.WithResponseSchema(schema.APIResponse{}),
	)

	//rootRouter.Post("/crud/user", userHandler.CreateUser,
	//	autofiber.WithDescription("Create new user"),
	//	autofiber.WithTags("admin", "user"),
	//	autofiber.WithRequestSchema(schema.CreateUserRequest{}),
	//	autofiber.WithResponseSchema(schema.APIResponse{}),
	//)

	//rootRouter.Put("/crud/user/:id", userHandler.UpdateUser,
	//	autofiber.WithDescription("Update user by ID"),
	//	autofiber.WithTags("admin", "user"),
	//	autofiber.WithRequestSchema(schema.UpdateUserRequest{}),
	//	autofiber.WithResponseSchema(schema.APIResponse{}),
	//)

	//rootRouter.Delete("/crud/user/:id", userHandler.DeleteUser,
	//	autofiber.WithDescription("Delete user by ID"),
	//	autofiber.WithTags("admin", "user"),
	//	autofiber.WithRequestSchema(schema.UserIDParam{}),
	//	autofiber.WithResponseSchema(schema.APIResponse{}),
	//)
}
