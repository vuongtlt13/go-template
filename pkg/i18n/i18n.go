package i18n

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"yourapp/pkg/config"
)

var (
	DefaultLocale    string
	SupportedLocales []string
	translations     = make(map[string]map[string]interface{})
	mutex            sync.RWMutex
)

// Init initializes the i18n package
func Init(cfg *config.I18nConfig) {
	DefaultLocale = cfg.DefaultLocale

	// Dynamically detect supported locales
	SupportedLocales = detectSupportedLocales(cfg.BaseFolder)

	// Load translations from files
	for _, locale := range SupportedLocales {
		loadTranslations(locale, cfg.BaseFolder)
	}
}

// detectSupportedLocales scans the base folder for locale directories
func detectSupportedLocales(baseFolder string) []string {
	entries, err := os.ReadDir(baseFolder)
	if err != nil {
		panic(err)
	}
	var locales []string
	for _, entry := range entries {
		if entry.IsDir() {
			locales = append(locales, entry.Name())
		}
	}
	return locales
}

// Helper to recursively merge maps
func mergeMap(dst, src map[string]interface{}) {
	for k, v := range src {
		if vMap, ok := v.(map[string]interface{}); ok {
			if dstMap, ok := dst[k].(map[string]interface{}); ok {
				mergeMap(dstMap, vMap)
			} else {
				dst[k] = vMap
			}
		} else {
			dst[k] = v
		}
	}
}

// loadTranslations loads translations from JSON files in the locale directory
func loadTranslations(locale string, baseFolder string) {
	basePath := filepath.Join(baseFolder, locale)
	translations[locale] = make(map[string]interface{})

	err := filepath.Walk(basePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() || !strings.HasSuffix(info.Name(), ".json") {
			return nil
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		var trans map[string]interface{}
		if err := json.Unmarshal(data, &trans); err != nil {
			return err
		}
		relPath, err := filepath.Rel(basePath, path)
		if err != nil {
			return err
		}
		keyParts := strings.Split(strings.TrimSuffix(relPath, ".json"), string(os.PathSeparator))
		current := translations[locale]
		for i, part := range keyParts {
			if i == len(keyParts)-1 {
				// Merge the loaded map here
				if existing, ok := current[part].(map[string]interface{}); ok {
					mergeMap(existing, trans)
				} else {
					current[part] = trans
				}
			} else {
				if _, ok := current[part]; !ok {
					current[part] = make(map[string]interface{})
				}
				current = current[part].(map[string]interface{})
			}
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}

// IsSupportedLocale checks if the given locale is supported
func IsSupportedLocale(locale string) bool {
	for _, l := range SupportedLocales {
		if l == locale {
			return true
		}
	}
	return false
}

// getNestedValue gets a value from a nested map using dot notation
func getNestedValue(m map[string]interface{}, key string) (interface{}, bool) {
	parts := strings.Split(key, ".")
	current := m

	for i, part := range parts {
		if i == len(parts)-1 {
			// Last part, return the value
			if val, ok := current[part]; ok {
				return val, true
			}
			return nil, false
		}

		// Navigate through the map
		if next, ok := current[part].(map[string]interface{}); ok {
			current = next
		} else {
			return nil, false
		}
	}

	return nil, false
}

// T translates a key to the given locale
func T(ctx context.Context, key string) string {
	// Get locale from context
	locale, ok := ctx.Value("lang").(string)
	if !ok {
		locale = DefaultLocale
	}

	// Get translation
	mutex.RLock()
	defer mutex.RUnlock()

	// Try to get translation from the specified locale
	if trans, ok := translations[locale]; ok {
		if val, ok := getNestedValue(trans, key); ok {
			if str, ok := val.(string); ok {
				return str
			}
		}
	}

	// Fallback to default locale
	if trans, ok := translations[DefaultLocale]; ok {
		if val, ok := getNestedValue(trans, key); ok {
			if str, ok := val.(string); ok {
				return str
			}
		}
	}

	// Return key if no translation found
	return key
}

// AddTranslation adds a translation
func AddTranslation(locale string, key string, message string) {
	mutex.Lock()
	defer mutex.Unlock()

	if _, ok := translations[locale]; !ok {
		translations[locale] = make(map[string]interface{})
	}

	parts := strings.Split(key, ".")
	current := translations[locale]

	for i, part := range parts {
		if i == len(parts)-1 {
			// Last part, set the value
			current[part] = message
			break
		}

		// Create nested map if it doesn't exist
		if _, ok := current[part]; !ok {
			current[part] = make(map[string]interface{})
		}

		// Navigate to the next level
		if next, ok := current[part].(map[string]interface{}); ok {
			current = next
		} else {
			// If the current value is not a map, replace it with a new map
			current[part] = make(map[string]interface{})
			current = current[part].(map[string]interface{})
		}
	}
}

// GetTranslations returns all translations for a specific language
func GetTranslations(locale string) map[string]interface{} {
	mutex.RLock()
	defer mutex.RUnlock()

	// If locale is not supported, use default locale
	if !IsSupportedLocale(locale) {
		locale = DefaultLocale
	}

	// Return a copy of the translations to prevent modification
	translations := make(map[string]interface{})
	if trans, ok := translations[locale].(map[string]interface{}); ok {
		for k, v := range trans {
			translations[k] = v
		}
	}

	return translations
}
