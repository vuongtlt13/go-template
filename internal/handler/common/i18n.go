package common

import (
	"yourapp/internal/schema"
	"yourapp/pkg/i18n"
	"yourapp/pkg/response"

	"github.com/gofiber/fiber/v2"
)

type I18nHandler struct {
}

func NewI18nHandler() *I18nHandler {
	return &I18nHandler{}
}

// GetTranslations returns all translations for a specific language
func (h *I18nHandler) GetTranslations(c *fiber.Ctx, req *schema.LanguageParam) (*response.APIResponse[map[string]interface{}], error) {
	lang := req.Lang

	// Validate language
	if !i18n.IsSupportedLocale(lang) {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Unsupported language")
	}

	// Get translations for the language
	translations := i18n.GetTranslations(lang)

	return &response.APIResponse[map[string]interface{}]{
		Success: true,
		Data:    translations,
		Message: "OK",
	}, nil
}
