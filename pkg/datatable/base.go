package datatable

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"yourapp/pkg/httperror"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

const DefaultMaxLimit = 300
const DefaultLimit = 25

// DataTable defines the interface for all datatables
// All datatable implementations should satisfy this interface
type DataTable interface {
	Render(c *fiber.Ctx, extra map[string]interface{}) error

	AddColumn(columnName string, producer ProducerFunc) DataTable
	EditColumn(columnName string, producer ProducerFunc) DataTable
	FilterColumn(columnName string, filter FilterFunc) DataTable
	AddIndexColumn(columnName string) DataTable
}

type OverrideDataTable interface {
	GetQuery() *gorm.DB
	GetColumns() []*ColumnDefinition
	ModifyDatatable()
	BeforeProcess()
	AfterProcess()
}

type DataTaleConfig struct {
	MaxLimit    int
	SmartSearch bool
}

type BaseDataTaleAction struct {
	GetQuery        func() *gorm.DB
	GetColumns      func() []*ColumnDefinition
	ModifyDatatable ModifyDatatableFunc
	BeforeProcess   BeforeProcessFunc
	AfterProcess    AfterProcessFunc
}

// BaseDataTable represents the base datatable with common logic and abstract methods
type BaseDataTable struct {
	fiberCtx          *fiber.Ctx
	db                *gorm.DB
	Query             *gorm.DB
	Columns           []*ColumnDefinition
	columnDefinitions map[string]*ColumnDefinition

	// Advanced features
	cfg      *DataTaleConfig
	override OverrideDataTable
	option   *DataTaleOption

	additionalCols  map[string]ProducerFunc
	editCols        map[string]ProducerFunc
	includeIndex    bool
	indexColumnName string
	filterRules     map[string]FilterFunc

	// State
	prepared        bool
	Result          any
	totalRecords    int64
	filteredRecords int64
}

func NewBaseDataTable(cfg *DataTaleConfig, db *gorm.DB) *BaseDataTable {
	dt := &BaseDataTable{
		db:                db,
		fiberCtx:          nil,
		Query:             nil,
		Columns:           nil,
		columnDefinitions: make(map[string]*ColumnDefinition),
		cfg:               cfg,
		option:            nil,
		additionalCols:    make(map[string]ProducerFunc),
		editCols:          make(map[string]ProducerFunc),
		filterRules:       make(map[string]FilterFunc),
		includeIndex:      false,
		indexColumnName:   "DT_RowIndex",
	}
	return dt
}

// GetDB returns the database connection
func (bdt *BaseDataTable) GetDB() *gorm.DB {
	return bdt.db
}

func (bdt *BaseDataTable) SetOverride(override OverrideDataTable) {
	bdt.override = override
	bdt.Columns = bdt.override.GetColumns()
	bdt.override.ModifyDatatable()
}

// Render processes the datatable based on action type
// Similar to FastAPI's render method
func (bdt *BaseDataTable) Render(c *fiber.Ctx, extra map[string]interface{}) (interface{}, error) {
	bdt.fiberCtx = c

	// Parse request
	req, err := ParseDataTableOptionFromFiberContext(c)
	if err != nil {
		return nil, httperror.NewBadRequest("Invalid request format")
	}

	bdt.option = req
	bdt.Query = bdt.override.GetQuery()
	bdt.override.BeforeProcess()

	// Call action based on request action
	var result interface{}
	var err2 error

	switch req.Action {
	case ActionAJAX:
		result, err2 = bdt.callAJAX(extra)
	case ActionExcel:
		result, err2 = bdt.callExcel()
	case ActionCSV:
		result, err2 = bdt.callCSV()
	case ActionPDF:
		result, err2 = bdt.callPDF()
	default:
		result, err2 = bdt.callAJAX(extra)
	}

	if err2 != nil {
		return nil, httperror.NewBadRequest(err2.Error())
	}
	bdt.override.AfterProcess()

	// Send response based on action type
	switch req.Action {
	case ActionExcel, ActionCSV, ActionPDF:
		// For export actions, send file response
		//return bdt.sendFileResponse(c, result, req.Action)
		return nil, nil
	default:
		// For AJAX, send JSON response
		if resp, ok := result.(*Result); ok {
			return resp, nil
		}
		return nil, httperror.NewBadRequest("Invalid response type")
	}
}

func (bdt *BaseDataTable) getSearchableColumnNames() []*ColumnDefinition {
	var result []*ColumnDefinition
	for _, col := range bdt.Columns {
		if col.Searchable {
			result = append(result, col)
		}
	}
	return result
}

func (bdt *BaseDataTable) getOrderableColumnNames() []*ColumnDefinition {
	var result []*ColumnDefinition
	for _, col := range bdt.Columns {
		if col.Orderable && col.ColumnAlias != "" {
			result = append(result, col)
		}
	}
	return result
}

func (bdt *BaseDataTable) getFilterColumns() []*ColumnDefinition {
	var result []*ColumnDefinition
	for _, col := range bdt.Columns {
		if col.ColumnAlias != "" {
			result = append(result, col)
		}
	}
	return result
}

func (bdt *BaseDataTable) countTotal(baseQuery *gorm.DB) (int64, error) {
	var result int64
	tx := bdt.db.Table("(?) as base_tbl", baseQuery).Count(&result)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return result, nil
}

func (bdt *BaseDataTable) filterSelectedIDs(ids []int) *gorm.DB {
	return bdt.Query.Where("id IN ?", ids)
}

func (bdt *BaseDataTable) extractFilterConfig(filterColumn *ColumnDefinition) (filterRule string, filterValue interface{}) {
	filterRule = ""
	filterValue = nil

	// Lấy filter config từ query param theo alias
	filterConfig := bdt.fiberCtx.Query(filterColumn.ColumnAlias, "")
	if filterConfig != "" && strings.HasPrefix(filterConfig, "::") {
		// Dùng regex để tách rule và value
		re := regexp.MustCompile(`::([^(]+)\((.*)\)`)
		matches := re.FindStringSubmatch(filterConfig)
		if len(matches) == 3 {
			filterRule = matches[1]
			// Parse value từ json
			if err := json.Unmarshal([]byte(matches[2]), &filterValue); err != nil {
				// Có thể log lỗi nếu muốn
			}
		}
	}

	return filterRule, filterValue
}

func (bdt *BaseDataTable) applyFilterRules() *gorm.DB {
	filterColumns := bdt.getFilterColumns()
	db := bdt.Query

	for _, fc := range filterColumns {
		rule, value := bdt.extractFilterConfig(fc)
		ruleStr := fmt.Sprintf("%v", rule) // Ép về string để so sánh

		colName := fc.Data

		switch ruleStr {
		case string(FilterRuleEqual):
			db = db.Where(fmt.Sprintf("%s = ?", colName), value)
		case string(FilterRuleLessThan):
			db = db.Where(fmt.Sprintf("%s < ?", colName), value)
		case string(FilterRuleLessThanEqual):
			db = db.Where(fmt.Sprintf("%s <= ?", colName), value)
		case string(FilterRuleGreaterThan):
			db = db.Where(fmt.Sprintf("%s > ?", colName), value)
		case string(FilterRuleGreaterThanEqual):
			db = db.Where(fmt.Sprintf("%s >= ?", colName), value)
		case string(FilterRuleBetween):
			if vals, ok := value.([]interface{}); ok && len(vals) == 2 {
				db = db.Where(fmt.Sprintf("%s BETWEEN ? AND ?", colName), vals[0], vals[1])
			}
		case string(FilterRuleIn):
			db = db.Where(fmt.Sprintf("%s IN ?", colName), value)
		case string(FilterRuleNotIn):
			db = db.Where(fmt.Sprintf("%s NOT IN ?", colName), value)
		default:
			// Có thể log hoặc bỏ qua
		}
	}

	return db
}

func (bdt *BaseDataTable) doSearch(keywords []string) *gorm.DB {
	// Loại bỏ trùng lặp
	unique := make(map[string]struct{})
	var filtered []string
	for _, k := range keywords {
		if k == "" {
			continue
		}
		if _, ok := unique[k]; !ok {
			unique[k] = struct{}{}
			filtered = append(filtered, k)
		}
	}
	keywords = filtered

	searchableCols := bdt.getSearchableColumnNames() // Trả về []string tên cột
	if len(searchableCols) == 0 || len(keywords) == 0 {
		return bdt.Query
	}

	// Xây dựng điều kiện OR cho từng keyword trên từng cột
	db := bdt.Query
	for _, keyword := range keywords {
		conds := make([]string, 0)
		args := make([]interface{}, 0)
		for _, col := range searchableCols {
			conds = append(conds, fmt.Sprintf("%s ILIKE ?", col.Data))
			args = append(args, "%"+fmt.Sprintf("%s", keyword)+"%")
		}
		// Gộp các điều kiện OR cho mỗi keyword
		db = db.Where("("+strings.Join(conds, " OR ")+")", args...)
	}

	// Nếu có filter rules custom, gọi tiếp hàm applyFilterRules nếu bạn có
	bdt.applyFilterRules()

	return db
}

func (bdt *BaseDataTable) doSmartSearch() *gorm.DB {
	keywords := strings.Fields(bdt.option.Keyword)
	return bdt.doSearch(keywords)
}

func (bdt *BaseDataTable) applyFilterRecords() *gorm.DB {
	// Nếu có selected_ids, filter theo selected_ids
	if len(bdt.option.SelectedIDs) > 0 {
		return bdt.filterSelectedIDs(bdt.option.SelectedIDs)
	}

	// Nếu smart_search, thực hiện smart search
	if bdt.cfg.SmartSearch {
		return bdt.doSmartSearch()
	}

	// Ngược lại, search theo keyword thường
	return bdt.doSearch([]string{bdt.option.Keyword})
}

func (bdt *BaseDataTable) prepareQuery(paginate bool) error {
	if !bdt.prepared {
		totalRecords, err := bdt.countTotal(bdt.Query)
		if err != nil {
			return fmt.Errorf("error when counting record %v", err)
		}
		bdt.totalRecords = totalRecords

		if bdt.totalRecords > 0 {
			bdt.applyFilterRecords()
			//bdt.applyOrder()
			if paginate {
				//bdt.applyPaginate()
			}
		}
		bdt.prepared = true
	}

	return nil
}

func (bdt *BaseDataTable) getResultQuery() error {
	filteredRecords, err := bdt.countTotal(bdt.Query)
	if err != nil {
		return fmt.Errorf("error when counting record %v", err)
	}
	bdt.filteredRecords = filteredRecords

	tx := bdt.Query.Find(&bdt.Result)
	if tx.Error != nil {
		return fmt.Errorf("error when finding record %v", err)
	}
	return nil
}

// ProcessWithBase processes the datatable with base configuration
func (bdt *BaseDataTable) ProcessWithBase(paginate bool) (any, error) {
	err := bdt.prepareQuery(paginate)
	if err != nil {
		return nil, fmt.Errorf("error when preparing query %v", err)
	}

	if err = bdt.getResultQuery(); err != nil {
		return nil, fmt.Errorf("error when getting query result %v", err)
	}

	return bdt.Result, nil
}

// callAJAX handles AJAX action
func (bdt *BaseDataTable) callAJAX(extra map[string]interface{}) (*Result, error) {
	result, err := bdt.ProcessWithBase(true)
	if err != nil {
		return nil, fmt.Errorf("error when proccessing ajax %v", err)
	}

	if extra == nil {
		extra = map[string]interface{}{}
	}

	return &Result{
		TotalRecords:    bdt.totalRecords,
		FilteredRecords: bdt.filteredRecords,
		Items:           result,
		Others:          extra,
	}, nil
}

// callExcel handles Excel export action
func (bdt *BaseDataTable) callExcel() (any, error) {
	result, err := bdt.ProcessWithBase(false)
	if err != nil {
		return nil, fmt.Errorf("error when exporting excel %v", err)
	}

	return result, nil
}

// callCSV handles CSV export action
func (bdt *BaseDataTable) callCSV() (interface{}, error) {
	return nil, fmt.Errorf("export csv is not supported")
}

// callPDF handles PDF export action
func (bdt *BaseDataTable) callPDF() (interface{}, error) {
	return nil, fmt.Errorf("export pdf is not supported")
}

// sendFileResponse sends file response for export actions
func (bdt *BaseDataTable) sendFileResponse(c *fiber.Ctx, result interface{}, action Action) error {
	// TODO: Implement file response sending
	// This would set appropriate headers and send file data
	switch action {
	case ActionExcel:
		c.Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
		c.Set("Content-Disposition", "attachment; filename=export.xlsx")
	case ActionCSV:
		c.Set("Content-Type", "text/csv")
		c.Set("Content-Disposition", "attachment; filename=export.csv")
	case ActionPDF:
		c.Set("Content-Type", "application/pdf")
		c.Set("Content-Disposition", "attachment; filename=export.pdf")
	}

	return c.JSON(result)
}
