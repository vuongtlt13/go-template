package i18n

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"yourapp/pkg/config"
)

func setupTestTranslations(t *testing.T) {
	// Create a temporary directory for test files
	tmpDir := t.TempDir()

	// Create test translation files
	createTestTranslationFiles(t, tmpDir)

	// Initialize i18n with the test directory
	Init(&config.I18nConfig{
		DefaultLocale: "en",
		BaseFolder:    tmpDir,
	})
}

func TestI18n(t *testing.T) {
	setupTestTranslations(t)

	tests := []struct {
		name     string
		ctx      context.Context
		key      string
		expected string
	}{
		{
			name:     "English user model translation",
			ctx:      context.WithValue(context.Background(), "lang", "en"),
			key:      "models.user.menu_title",
			expected: "User",
		},
		{
			name:     "Vietnamese user model translation",
			ctx:      context.WithValue(context.Background(), "lang", "vi"),
			key:      "models.user.menu_title",
			expected: "Người dùng",
		},
		{
			name:     "English auth translation",
			ctx:      context.WithValue(context.Background(), "lang", "en"),
			key:      "auth.login.title",
			expected: "Login",
		},
		{
			name:     "Vietnamese auth translation",
			ctx:      context.WithValue(context.Background(), "lang", "vi"),
			key:      "auth.login.title",
			expected: "Đăng nhập",
		},
		{
			name:     "English CRUD action translation",
			ctx:      context.WithValue(context.Background(), "lang", "en"),
			key:      "crud.actions.create",
			expected: "Create",
		},
		{
			name:     "Vietnamese CRUD action translation",
			ctx:      context.WithValue(context.Background(), "lang", "vi"),
			key:      "crud.actions.create",
			expected: "Tạo mới",
		},
		{
			name:     "Fallback to default locale",
			ctx:      context.WithValue(context.Background(), "lang", "fr"),
			key:      "models.user.menu_title",
			expected: "User",
		},
		{
			name:     "Missing translation returns key",
			ctx:      context.WithValue(context.Background(), "lang", "en"),
			key:      "nonexistent.key",
			expected: "nonexistent.key",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := T(tt.ctx, tt.key)
			if result != tt.expected {
				t.Errorf("T() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestIsSupportedLocale(t *testing.T) {
	setupTestTranslations(t)

	tests := []struct {
		name     string
		locale   string
		expected bool
	}{
		{
			name:     "Supported English locale",
			locale:   "en",
			expected: true,
		},
		{
			name:     "Supported Vietnamese locale",
			locale:   "vi",
			expected: true,
		},
		{
			name:     "Unsupported locale",
			locale:   "fr",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsSupportedLocale(tt.locale)
			if result != tt.expected {
				t.Errorf("IsSupportedLocale() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestNestedTranslations(t *testing.T) {
	setupTestTranslations(t)

	tests := []struct {
		name     string
		ctx      context.Context
		key      string
		expected string
	}{
		{
			name:     "Nested user field translation - English",
			ctx:      context.WithValue(context.Background(), "lang", "en"),
			key:      "models.user.fields.email",
			expected: "Email",
		},
		{
			name:     "Nested user field translation - Vietnamese",
			ctx:      context.WithValue(context.Background(), "lang", "vi"),
			key:      "models.user.fields.email",
			expected: "Email",
		},
		{
			name:     "Nested auth message translation - English",
			ctx:      context.WithValue(context.Background(), "lang", "en"),
			key:      "auth.login.invalid_credentials",
			expected: "Invalid email or password",
		},
		{
			name:     "Nested auth message translation - Vietnamese",
			ctx:      context.WithValue(context.Background(), "lang", "vi"),
			key:      "auth.login.invalid_credentials",
			expected: "Email hoặc mật khẩu không đúng",
		},
		{
			name:     "Nested CRUD message translation - English",
			ctx:      context.WithValue(context.Background(), "lang", "en"),
			key:      "crud.messages.create_success",
			expected: "Created successfully",
		},
		{
			name:     "Nested CRUD message translation - Vietnamese",
			ctx:      context.WithValue(context.Background(), "lang", "vi"),
			key:      "crud.messages.create_success",
			expected: "Tạo mới thành công",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := T(tt.ctx, tt.key)
			if result != tt.expected {
				t.Errorf("T() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestLoadTranslationsFromFolders(t *testing.T) {
	setupTestTranslations(t)

	tests := []struct {
		locale   string
		key      string
		expected string
	}{
		// Root-level file
		{"en", "auth.login.title", "Login"},
		{"vi", "auth.login.title", "Đăng nhập"},
		// Nested file (subfolder)
		{"en", "models.user.menu_title", "User"},
		{"vi", "models.user.menu_title", "Người dùng"},
		// Nested field in subfolder
		{"en", "models.user.fields.email", "Email"},
		{"vi", "models.user.fields.email", "Email"},
		// CRUD from root
		{"en", "crud.actions.create", "Create"},
		{"vi", "crud.actions.create", "Tạo mới"},
	}

	for _, tt := range tests {
		ctx := context.WithValue(context.Background(), "lang", tt.locale)
		got := T(ctx, tt.key)
		if got != tt.expected {
			t.Errorf("T(%q, %q) = %q, want %q", tt.locale, tt.key, got, tt.expected)
		}
	}
}

// createTestTranslationFiles creates test translation files in the given directory
func createTestTranslationFiles(t *testing.T, baseDir string) {
	// Create directory structure
	dirs := []string{
		filepath.Join(baseDir, "en", "models"),
		filepath.Join(baseDir, "vi", "models"),
	}
	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			t.Fatal(err)
		}
	}

	// English files
	enAuth := map[string]interface{}{
		"login": map[string]interface{}{
			"title":               "Login",
			"invalid_credentials": "Invalid email or password",
		},
	}
	enCrud := map[string]interface{}{
		"actions": map[string]interface{}{
			"create": "Create",
		},
		"messages": map[string]interface{}{
			"create_success": "Created successfully",
		},
	}
	enUser := map[string]interface{}{
		"menu_title": "User",
		"fields": map[string]interface{}{
			"email": "Email",
		},
	}

	// Vietnamese files
	viAuth := map[string]interface{}{
		"login": map[string]interface{}{
			"title":               "Đăng nhập",
			"invalid_credentials": "Email hoặc mật khẩu không đúng",
		},
	}
	viCrud := map[string]interface{}{
		"actions": map[string]interface{}{
			"create": "Tạo mới",
		},
		"messages": map[string]interface{}{
			"create_success": "Tạo mới thành công",
		},
	}
	viUser := map[string]interface{}{
		"menu_title": "Người dùng",
		"fields": map[string]interface{}{
			"email": "Email",
		},
	}

	// Write English files
	writeJSON(t, filepath.Join(baseDir, "en", "auth.json"), enAuth)
	writeJSON(t, filepath.Join(baseDir, "en", "crud.json"), enCrud)
	writeJSON(t, filepath.Join(baseDir, "en", "models", "user.json"), enUser)

	// Write Vietnamese files
	writeJSON(t, filepath.Join(baseDir, "vi", "auth.json"), viAuth)
	writeJSON(t, filepath.Join(baseDir, "vi", "crud.json"), viCrud)
	writeJSON(t, filepath.Join(baseDir, "vi", "models", "user.json"), viUser)
}

func writeJSON(t *testing.T, path string, data interface{}) {
	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, b, 0644); err != nil {
		t.Fatal(err)
	}
}
