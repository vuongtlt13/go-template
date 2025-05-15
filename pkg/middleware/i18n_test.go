package middleware

import (
	"encoding/json"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"yourapp/pkg/config"
	"yourapp/pkg/i18n"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestI18nMiddleware(t *testing.T) {
	// Create temporary directory for test translations
	tmpDir := t.TempDir()

	// Create test translation files
	createTestTranslationFiles(t, tmpDir)

	// Initialize i18n with test translations
	i18n.Init(&config.I18nConfig{
		DefaultLocale: "en",
		BaseFolder:    tmpDir,
	})

	app := fiber.New()

	app.Use(I18nMiddleware)

	// Add a test route to check the language in context
	app.Get("/", func(c *fiber.Ctx) error {
		lang := c.Locals("lang")
		return c.JSON(fiber.Map{"lang": lang})
	})

	tests := []struct {
		name           string
		acceptLanguage string
		expectedLang   string
	}{
		{
			name:           "English language",
			acceptLanguage: "en",
			expectedLang:   "en",
		},
		{
			name:           "Vietnamese language",
			acceptLanguage: "vi",
			expectedLang:   "vi",
		},
		{
			name:           "Multiple languages with English first",
			acceptLanguage: "en,vi;q=0.9",
			expectedLang:   "en",
		},
		{
			name:           "Multiple languages with Vietnamese first",
			acceptLanguage: "vi,en;q=0.9",
			expectedLang:   "vi",
		},
		{
			name:           "Unsupported language falls back to default",
			acceptLanguage: "fr",
			expectedLang:   "en",
		},
		{
			name:           "Empty header falls back to default",
			acceptLanguage: "",
			expectedLang:   "en",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/", nil)
			if tt.acceptLanguage != "" {
				req.Header.Set("Accept-Language", tt.acceptLanguage)
			}

			resp, err := app.Test(req)
			assert.NoError(t, err)
			assert.Equal(t, fiber.StatusOK, resp.StatusCode)

			// Parse response body
			var result fiber.Map
			err = json.NewDecoder(resp.Body).Decode(&result)
			assert.NoError(t, err)

			// Check language in response
			assert.Equal(t, tt.expectedLang, result["lang"])
		})
	}
}

// createTestTranslationFiles creates test translation files in the given directory
func createTestTranslationFiles(t *testing.T, baseDir string) {
	// Create English translations
	enDir := filepath.Join(baseDir, "en")
	err := os.MkdirAll(enDir, 0755)
	assert.NoError(t, err)

	// Create Vietnamese translations
	viDir := filepath.Join(baseDir, "vi")
	err = os.MkdirAll(viDir, 0755)
	assert.NoError(t, err)

	// Create test translation files
	translations := map[string]map[string]string{
		"en": {
			"auth.login.title":            "Login",
			"auth.login.submit":           "Sign In",
			"auth.login.success":          "Login successful",
			"auth.login.error":            "Invalid credentials",
			"crud.create":                 "Create",
			"crud.read":                   "Read",
			"crud.update":                 "Update",
			"crud.delete":                 "Delete",
			"models.user.menu_title":      "User",
			"models.user.fields.email":    "Email",
			"models.user.fields.password": "Password",
		},
		"vi": {
			"auth.login.title":            "Đăng nhập",
			"auth.login.submit":           "Đăng nhập",
			"auth.login.success":          "Đăng nhập thành công",
			"auth.login.error":            "Thông tin đăng nhập không hợp lệ",
			"crud.create":                 "Tạo mới",
			"crud.read":                   "Xem",
			"crud.update":                 "Cập nhật",
			"crud.delete":                 "Xóa",
			"models.user.menu_title":      "Người dùng",
			"models.user.fields.email":    "Email",
			"models.user.fields.password": "Mật khẩu",
		},
	}

	for locale, trans := range translations {
		// Create auth.json
		authData := map[string]interface{}{
			"login": map[string]interface{}{
				"title":   trans["auth.login.title"],
				"submit":  trans["auth.login.submit"],
				"success": trans["auth.login.success"],
				"error":   trans["auth.login.error"],
			},
		}
		authFile := filepath.Join(baseDir, locale, "auth.json")
		err = os.WriteFile(authFile, mustMarshalJSON(authData), 0644)
		assert.NoError(t, err)

		// Create crud.json
		crudData := map[string]interface{}{
			"create": trans["crud.create"],
			"read":   trans["crud.read"],
			"update": trans["crud.update"],
			"delete": trans["crud.delete"],
		}
		crudFile := filepath.Join(baseDir, locale, "crud.json")
		err = os.WriteFile(crudFile, mustMarshalJSON(crudData), 0644)
		assert.NoError(t, err)

		// Create models/user.json
		modelsDir := filepath.Join(baseDir, locale, "models")
		err = os.MkdirAll(modelsDir, 0755)
		assert.NoError(t, err)

		userData := map[string]interface{}{
			"menu_title": trans["models.user.menu_title"],
			"fields": map[string]interface{}{
				"email":    trans["models.user.fields.email"],
				"password": trans["models.user.fields.password"],
			},
		}
		userFile := filepath.Join(modelsDir, "user.json")
		err = os.WriteFile(userFile, mustMarshalJSON(userData), 0644)
		assert.NoError(t, err)
	}
}

// mustMarshalJSON marshals a value to JSON and panics on error
func mustMarshalJSON(v interface{}) []byte {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		panic(err)
	}
	return data
}
