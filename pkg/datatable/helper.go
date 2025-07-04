package datatable

import (
	"github.com/gofiber/fiber/v2"
)

// RequestParams represents common datatable parameters
type RequestParams struct {
	Keyword     string `parse:"query:q" validate:"omitempty" default:"" description:"Search keyword"`
	Skip        int    `parse:"query:s" validate:"omitempty,min=0" default:"0" description:"Offset"`
	Limit       int    `parse:"query:ipp" validate:"omitempty,min=1,max=100" default:"25" description:"Number of items per page"`
	Action      string `parse:"query:action" validate:"omitempty,max=255" default:"ajax" description:"Action type"`
	SelectedIDs []int  `parse:"query:ids" validate:"omitempty,max=255" description:"Selected ids"`
}

type IDatatableRequest interface {
	GetRequestParams() *RequestParams
}

// ParseDataTableOptionFromFiberContext parses datatable request from fiber context (GET only, new DataTaleOption structure)
func ParseDataTableOptionFromFiberContext(c *fiber.Ctx, reqParams IDatatableRequest) (*DataTaleOption, error) {
	var req DataTaleOption

	req.RawRequest = reqParams

	// Parse keyword (for smart search)
	datatableReq := reqParams.GetRequestParams()
	req.Keyword = datatableReq.Keyword
	req.Skip = datatableReq.Skip
	req.Limit = datatableReq.Limit
	if req.Limit == 0 {
		req.Limit = DefaultLimit
	}

	// Parse action
	req.Action = Action(datatableReq.Action)

	req.SelectedIDs = datatableReq.SelectedIDs

	return &req, nil
}
