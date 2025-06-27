package common

import (
	"yourapp/internal/handler/common"
	"yourapp/internal/schema"

	autofiber "github.com/vuongtlt13/auto-fiber"
)

type I18nRouter struct{}

func NewI18nRouter() *I18nRouter { return &I18nRouter{} }

func (r *I18nRouter) Register(rootRouter *autofiber.AutoFiberGroup) {
	handler := common.NewI18nHandler()
	rootRouter.Get("/i18n", handler.GetI18n,
		autofiber.WithDescription("Get internationalization data"),
		autofiber.WithTags("i18n"),
		autofiber.WithResponseSchema(schema.APIResponse{}),
	)

	rootRouter.Get("/i18n/:lang", handler.GetTranslations,
		autofiber.WithDescription("Get translations for specific language"),
		autofiber.WithTags("i18n"),
		autofiber.WithRequestSchema(schema.LanguageParam{}),
		autofiber.WithResponseSchema(schema.APIResponse{}),
	)
}
