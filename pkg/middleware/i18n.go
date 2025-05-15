package middleware

import (
	"strings"
	"yourapp/pkg/i18n"

	"github.com/gofiber/fiber/v2"
)

// I18nMiddleware is a middleware that handles internationalization
func I18nMiddleware(c *fiber.Ctx) error {
	// Get language from header
	lang := GetLangFromHeader(c.GetReqHeaders())

	// Set language in context
	c.Locals("lang", lang)

	return c.Next()
}

// GetLangFromHeader gets the language from the Accept-Language header
func GetLangFromHeader(header map[string][]string) string {
	// Get Accept-Language header
	acceptLang := header["Accept-Language"]
	if len(acceptLang) == 0 {
		return i18n.DefaultLocale
	}

	// Parse Accept-Language header
	// Example: "en-US,en;q=0.9,vi;q=0.8"
	langs := strings.Split(acceptLang[0], ",")
	if len(langs) == 0 {
		return i18n.DefaultLocale
	}

	// Get first language
	// Example: "en-US" -> "en"
	lang := strings.Split(langs[0], "-")[0]

	// Check if language is supported
	if !i18n.IsSupportedLocale(lang) {
		return i18n.DefaultLocale
	}

	return lang
}

// GetLang gets the language from the context
func GetLang(c *fiber.Ctx) string {
	lang, ok := c.Locals("lang").(string)
	if !ok {
		return i18n.DefaultLocale
	}
	return lang
}
