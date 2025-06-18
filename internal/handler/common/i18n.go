package common

import (
	"encoding/json"
	"yourapp/pkg/i18n"

	"github.com/gofiber/fiber/v2"
)

type I18nHandler struct {
	i18n *i18n.I18n
}

func NewI18nHandler(i18n *i18n.I18n) *I18nHandler {
	return &I18nHandler{
		i18n: i18n,
	}
}

// GetTranslations returns all translations for a specific language
func (h *I18nHandler) GetTranslations(c *fiber.Ctx) error {
	lang := c.Params("lang")

	// Validate language
	if !i18n.IsSupportedLocale(lang) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Unsupported language",
		})
	}

	// Get translations for the language
	translations := i18n.GetTranslations(lang)

	// Marshal translations to JSON
	jsonBytes, err := json.Marshal(translations)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to marshal translations",
		})
	}

	return c.Status(fiber.StatusOK).Send(jsonBytes)
}

func (h *I18nHandler) RegisterRoutes(app *fiber.App) {
	app.Get("/api/v1/i18n/:lang", h.GetTranslations)
}
