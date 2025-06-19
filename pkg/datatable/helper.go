package datatable

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// ParseDataTableOptionFromFiberContext parses datatable request from fiber context (GET only, new DataTaleOption structure)
func ParseDataTableOptionFromFiberContext(c *fiber.Ctx) (*DataTaleOption, error) {
	var req DataTaleOption
	var err error

	// Parse keyword (for smart search)
	req.Keyword = c.Query("q", "")
	req.Skip, err = strconv.Atoi(c.Query("s", "0"))
	if err != nil {
		req.Skip = 0
	}
	req.Limit, err = strconv.Atoi(c.Query("ipp", fmt.Sprintf("%d", DefaultLimit)))
	if err != nil {
		req.Limit = DefaultLimit
	}

	// Parse action
	req.Action = Action(c.Query("action", string(ActionAJAX)))

	// Parse selected IDs (as JSON array in query param, e.g. ?selected_ids=[1,2,3])
	selectedIDs := c.Query("ids", "[]")
	if selectedIDs != "[]" && selectedIDs != "" {
		var ids []int
		if err := json.Unmarshal([]byte(selectedIDs), &ids); err == nil {
			req.SelectedIDs = ids
		}
	}

	return &req, nil
}
