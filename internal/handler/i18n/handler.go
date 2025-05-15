package i18n

import (
	"context"
	"encoding/json"

	pb "yourapp/pb/i18n"
	i18nconnect "yourapp/pb/i18n/i18nconnect"
	"yourapp/pkg/i18n"

	"connectrpc.com/connect"
)

// Handler implements the i18n service
type Handler struct {
	i18nconnect.UnimplementedI18NServiceHandler
}

// NewHandler creates a new i18n handler
func NewHandler() *Handler {
	return &Handler{}
}

// GetTranslations returns all translations for a specific language
func (h *Handler) GetTranslations(
	ctx context.Context,
	req *connect.Request[pb.GetTranslationsRequest],
) (*connect.Response[pb.GetTranslationsResponse], error) {
	lang := req.Msg.Language

	// Validate language
	if !i18n.IsSupportedLocale(lang) {
		return nil, connect.NewError(connect.CodeInvalidArgument, nil)
	}

	// Get translations for the language
	translations := i18n.GetTranslations(lang)

	// Marshal translations to JSON
	jsonBytes, err := json.Marshal(translations)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	// Create response
	res := connect.NewResponse(&pb.GetTranslationsResponse{
		Translations: jsonBytes,
	})

	return res, nil
}
