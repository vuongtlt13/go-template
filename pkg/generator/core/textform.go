package core

import "github.com/jinzhu/inflection"

// TODO: Implement Singularize and Pluralize helpers for English nouns
func Singularize(s string) string {
	return inflection.Singular(s)
}

func Pluralize(s string) string {
	return inflection.Plural(s)
}

// TextForm holds various forms of a base text (usually snake_case)
type TextForm struct {
	OriginText      string
	SnakeCase       string
	SnakeCasePlural string
	KebabCase       string
	KebabCasePlural string
	HumanCase       string
	HumanCasePlural string
	CamelCase       string
	CamelCasePlural string
}

func NewTextForm(snakeCaseText string) *TextForm {
	singular := Singularize(snakeCaseText)
	plural := Pluralize(snakeCaseText)
	return &TextForm{
		OriginText:      snakeCaseText,
		SnakeCase:       singular,
		SnakeCasePlural: plural,
		KebabCase:       KebabCase(singular),
		KebabCasePlural: KebabCase(plural),
		HumanCase:       HumanCase(singular),
		HumanCasePlural: HumanCase(plural),
		CamelCase:       CamelCase(singular),
		CamelCasePlural: CamelCase(plural),
	}
}

// ModelTextForm holds various forms for a model name (any string)
type ModelTextForm struct {
	TableName                string
	ModelNameSnakeCase       string
	ModelNameSnakeCasePlural string
	ModelNameKebabCase       string
	ModelNameKebabCasePlural string
	ModelName                string
	ModelNamePlural          string
	ModelNameCamelCase       string
	ModelNameCamelCasePlural string
	ModelNameHuman           string
	ModelNameHumanPlural     string
}

func NewModelTextForm(anyString string) *ModelTextForm {
	tableName := Pluralize(SnakeCase(anyString))
	singular := Singularize(tableName)
	plural := Pluralize(tableName)
	modelName := PascalCase(singular)
	modelNamePlural := Pluralize(modelName)
	return &ModelTextForm{
		TableName:                tableName,
		ModelNameSnakeCase:       singular,
		ModelNameSnakeCasePlural: plural,
		ModelNameKebabCase:       KebabCase(singular),
		ModelNameKebabCasePlural: KebabCase(plural),
		ModelName:                modelName,
		ModelNamePlural:          modelNamePlural,
		ModelNameCamelCase:       CamelCase(singular),
		ModelNameCamelCasePlural: CamelCase(plural),
		ModelNameHuman:           HumanCase(singular),
		ModelNameHumanPlural:     HumanCase(plural),
	}
}
