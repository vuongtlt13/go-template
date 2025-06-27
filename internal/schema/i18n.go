package schema

// LanguageParam represents path parameter for language
type LanguageParam struct {
	Lang string `parse:"path:lang" validate:"required,min=2,max=5" description:"Language code (e.g., en, vi)"`
}
