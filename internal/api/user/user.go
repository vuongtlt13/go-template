package user

import (
	"yourapp/internal/handler/user"
	"yourapp/pkg/response"
	
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
	userHandler := user.NewUserHandler()

	rootRouter.Get("/user/profile", userHandler.GetProfile,
		autofiber.WithDescription("Get current user profile"),
		autofiber.WithTags("user", "profile"),
		autofiber.WithResponseSchema(response.APIResponse[map[string]interface{}]{}),
	)

	// Thêm các route user khác nếu cần
}
