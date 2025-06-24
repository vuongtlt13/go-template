package datatable

import "gorm.io/gorm"

// Action represents the datatable action type
type Action string

const (
	ActionAJAX  Action = "ajax"
	ActionExcel Action = "excel"
	ActionCSV   Action = "csv"
	ActionPDF   Action = "pdf"
)

// FilterRuleType represents different filter rule types
type FilterRuleType string

const (
	FilterRuleLike             FilterRuleType = "like"
	FilterRuleEqual            FilterRuleType = "eq"
	FilterRuleGreaterThan      FilterRuleType = "gt"
	FilterRuleGreaterThanEqual FilterRuleType = "gte"
	FilterRuleLessThan         FilterRuleType = "lt"
	FilterRuleLessThanEqual    FilterRuleType = "lte"
	FilterRuleBetween          FilterRuleType = "bt"
	FilterRuleIn               FilterRuleType = "in"
	FilterRuleNotIn            FilterRuleType = "not_in"
)

// DataTaleOption represents the datatable request from frontend
type DataTaleOption struct {
	Keyword string `json:"keyword"`
	Skip    int    `json:"skip"`
	Limit   int    `json:"limit"`
	Action  Action `json:"action"`

	//SortBys   []string `json:"sort_bys"`
	//SortTypes []string `json:"sort_types"`
	Others map[string]interface{}

	SelectedIDs []int `json:"selected_ids,omitempty"`
}

// Column represents column information
type Column struct {
	Data       string `json:"data"`
	Name       string `json:"name"`
	Searchable bool   `json:"searchable"`
	Orderable  bool   `json:"orderable"`
	Exportable bool   `json:"exportable"`
	Printable  bool   `json:"printable"`
}

// Result represents the datatable response
type Result struct {
	TotalRecords    int64                  `json:"totalRecords"`
	FilteredRecords int64                  `json:"filteredRecords"`
	Items           any                    `json:"items,omitempty"`
	Others          map[string]interface{} `json:"others,omitempty"`
}

// ColumnDefinition represents a column definition
type ColumnDefinition struct {
	Data        string
	Title       string
	Searchable  bool
	Orderable   bool
	Exportable  bool
	Printable   bool
	ColumnAlias string
	Filter      FilterFunc
}

// FilterFunc represents a custom filter function
type FilterFunc func(value string) interface{}

type GetQueryFunc func() *gorm.DB
type GetColumnsFunc func() []*ColumnDefinition
type ModifyDatatableFunc func()
type Reducer func(rows []interface{}) []interface{}
type BeforeProcessFunc func()
type AfterProcessFunc func()

// ProducerFunc represents a function that produces additional column data
type ProducerFunc func(record interface{}) interface{}
