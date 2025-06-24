package core

import (
	"strings"

	"github.com/iancoleman/strcase"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// SnakeCase converts a string to snake_case
func SnakeCase(s string) string {
	return strcase.ToSnake(s)
}

// CamelCase converts a string to camelCase
func CamelCase(s string) string {
	return strcase.ToLowerCamel(s)
}

// PascalCase converts a string to PascalCase
func PascalCase(s string) string {
	return strcase.ToCamel(s)
}

// KebabCase converts a string to kebab-case
func KebabCase(s string) string {
	return strcase.ToKebab(s)
}

// HumanCase converts a string to Human Case
func HumanCase(s string) string {
	// Tách snake_case/kebab-case/camelCase thành các từ, rồi TitleCase từng từ
	s = strcase.ToDelimited(s, ' ')
	words := strings.Fields(s)
	titleCaser := cases.Title(language.Und, cases.NoLower)
	for i := range words {
		words[i] = titleCaser.String(words[i])
	}
	return strings.Join(words, " ")
}
