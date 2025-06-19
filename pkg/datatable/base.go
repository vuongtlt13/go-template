package datatable

import (
	"fmt"
	"yourapp/pkg/response"

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
	Reducer(rows []interface{}) []interface{}
}

type DataTaleConfig struct {
	maxLimit    int
	smartSearch bool
}

// BaseDataTable represents the base datatable with common logic and abstract methods
type BaseDataTable struct {
	fiberCtx          *fiber.Ctx
	db                *gorm.DB
	Query             *gorm.DB
	Columns           []*ColumnDefinition
	ReducerFunc       Reducer
	BeforeProcessFunc BeforeProcessFunc
	columnDefinitions map[string]*ColumnDefinition

	// Advanced features
	cfg    *DataTaleConfig
	option *DataTaleOption

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

func NewBaseDataTable(cfg *DataTaleConfig, db *gorm.DB) BaseDataTable {
	if cfg == nil {
		cfg = &DataTaleConfig{
			maxLimit:    DefaultMaxLimit,
			smartSearch: true,
		}
	}

	return BaseDataTable{
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
}

// GetDB returns the database connection
func (bdt *BaseDataTable) GetDB() *gorm.DB {
	return bdt.db
}

// Render processes the datatable based on action type
// Similar to FastAPI's render method
func (bdt *BaseDataTable) Render(c *fiber.Ctx, extra map[string]interface{}) error {
	bdt.fiberCtx = c

	// Parse request
	req, err := ParseDataTableOptionFromFiberContext(c)
	if err != nil {
		return response.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request format", fiber.StatusBadRequest)
	}

	bdt.option = req
	bdt.BeforeProcessFunc()

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
		return response.ErrorResponse(c, fiber.StatusBadRequest, err2.Error(), fiber.StatusBadRequest)
	}

	// Send response based on action type
	switch req.Action {
	case ActionExcel, ActionCSV, ActionPDF:
		// For export actions, send file response
		return bdt.sendFileResponse(c, result, req.Action)
	default:
		// For AJAX, send JSON response
		if resp, ok := result.(*Result); ok {
			return response.SuccessResponse(c, resp, "ok")
		}
		return response.ErrorResponse(c, fiber.StatusBadRequest, "Invalid response type", fiber.StatusBadRequest)
	}
}

// GetSearchableColumns returns the searchable columns
// Override this method in child classes
func (bdt *BaseDataTable) GetSearchableColumns() []string {
	return []string{}
}

// GetOrderableColumns returns the orderable columns
// Override this method in child classes
func (bdt *BaseDataTable) GetOrderableColumns() []string {
	return []string{}
}

// GetRelations returns the relations to preload
// Override this method in child classes
func (bdt *BaseDataTable) GetRelations() []string {
	return []string{}
}

// GetSelects returns the select fields
// Override this method in child classes
func (bdt *BaseDataTable) GetSelects() []string {
	return []string{}
}

// GetConditions returns the where conditions
// Override this method in child classes
func (bdt *BaseDataTable) GetConditions() map[string]interface{} {
	return make(map[string]interface{})
}

// GetAdditionalColumns returns additional columns configuration
// Override this method in child classes
func (bdt *BaseDataTable) GetAdditionalColumns() map[string]ProducerFunc {
	return make(map[string]ProducerFunc)
}

// GetEditColumns returns edit columns configuration
// Override this method in child classes
func (bdt *BaseDataTable) GetEditColumns() map[string]ProducerFunc {
	return make(map[string]ProducerFunc)
}

// GetFilterRules returns custom filter rules
// Override this method in child classes
func (bdt *BaseDataTable) GetFilterRules() map[string]FilterFunc {
	return make(map[string]FilterFunc)
}

// ShouldIncludeIndex returns whether to include index column
// Override this method in child classes
func (bdt *BaseDataTable) ShouldIncludeIndex() bool {
	return false
}

// GetIndexColumnName returns the index column name
// Override this method in child classes
func (bdt *BaseDataTable) GetIndexColumnName() string {
	return "DT_RowIndex"
}

// GetMaxLimit returns the maximum limit
// Override this method in child classes
func (bdt *BaseDataTable) GetMaxLimit() int {
	return 300
}

// ShouldUseSmartSearch returns whether to use smart search
// Override this method in child classes
func (bdt *BaseDataTable) ShouldUseSmartSearch() bool {
	return true
}

func (bdt *BaseDataTable) countTotal(baseQuery *gorm.DB) (int64, error) {
	var result int64
	tx := bdt.db.Table("(?) as base_tbl", baseQuery).Count(&result)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return result, nil
}

func (bdt *BaseDataTable) prepareQuery(paginate bool) error {
	if !bdt.prepared {
		totalRecords, err := bdt.countTotal(bdt.Query)
		if err != nil {
			return fmt.Errorf("error when counting record %v", err)
		}
		bdt.totalRecords = totalRecords

		if bdt.totalRecords > 0 {
			//bdt.applyFilterRecords()
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
