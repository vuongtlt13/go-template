package routes

import (
	"yourapp/internal/api/admin"
	"yourapp/internal/api/common"
	"yourapp/internal/repository"
	"yourapp/internal/service"
	"yourapp/pkg/config"
	"yourapp/pkg/database"

	autofiber "github.com/vuongtlt13/auto-fiber"
)

// Router represents admin routes
type AdminRouter struct {
	app *autofiber.AutoFiber
}

// NewRouter creates a new admin router
func NewAdminRouter() *AdminRouter {
	return &AdminRouter{
		app: nil,
	}
}

// Register registers all admin routes
func (r *AdminRouter) Register(app *autofiber.AutoFiber) {
	r.app = app
	api := r.app.Group("/api")

	// Khởi tạo config và authService tại đây
	cfg := config.GetConfig()
	db := database.GetDatabase()
	repo := repository.NewUserRepository(db)
	authService := service.NewAuthService(db, repo, nil) // truyền jwtManager nếu cần

	// Register common routes
	common.NewAuthRouter(cfg, authService).Register(api)
	common.NewI18nRouter().Register(api)
	common.NewProfileRouter().Register(api)
	common.NewHealthRouter().Register(api)

	// User management
	userRouter := admin.NewUserRouter()
	userRouter.Register(api)
}
