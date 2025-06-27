package routes

import (
	"yourapp/internal/api/common"
	"yourapp/internal/repository"
	"yourapp/internal/service"
	"yourapp/pkg/config"
	"yourapp/pkg/database"

	autofiber "github.com/vuongtlt13/auto-fiber"
)

// UserRouter represents user routes
type UserRouter struct {
	app *autofiber.AutoFiber
}

// NewUserRouter creates a new user router
func NewUserRouter() *UserRouter {
	return &UserRouter{app: nil}
}

// Register registers all user routes
func (r *UserRouter) Register(app *autofiber.AutoFiber) {
	r.app = app
	api := r.app.Group("/api")

	// Khởi tạo config và authService tại đây
	cfg := config.GetConfig()
	db := database.GetDatabase()
	repo := repository.NewUserRepository(db)
	authService := service.NewAuthService(db, repo, nil) // truyền jwtManager nếu cần

	// Register common routes
	common.NewHealthRouter().Register(api)
	common.NewAuthRouter(cfg, authService).Register(api)
	common.NewI18nRouter().Register(api)
	common.NewProfileRouter().Register(api)
}
