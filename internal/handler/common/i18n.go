package common

import (
	"encoding/json"
	"yourapp/pkg/i18n"

	"github.com/gofiber/fiber/v2"
)

type I18nHandler struct {
}

func NewI18nHandler() *I18nHandler {
	return &I18nHandler{}
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

// GetI18n returns i18n translations
// @Summary Get i18n translations
// @Description Get all translations for the current language
// @Tags i18n
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/i18n [get]
func (h *I18nHandler) GetI18n(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "i18n handler stub"})
}
