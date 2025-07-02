package common

import (
	"yourapp/internal/handler/common"

	autofiber "github.com/vuongtlt13/auto-fiber"
)

type ProfileRouter struct{}

func NewProfileRouter() *ProfileRouter { return &ProfileRouter{} }

func (r *ProfileRouter) Register(rootRouter *autofiber.AutoFiberGroup) {
	handler := common.NewProfileHandler()
	rootRouter.Get("/profile", handler.GetProfile,
		autofiber.WithDescription("Get user profile information"),
		autofiber.WithTags("profile"),
		autofiber.WithResponseSchema(response.APIResponse[map[string]interface{}]{}),
	)
}
