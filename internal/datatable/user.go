package datatable

import (
	"yourapp/internal/model"
	"yourapp/pkg/database"
	"yourapp/pkg/datatable"

	"gorm.io/gorm"
)

type UserDataTable struct {
	datatable.BaseDataTable
}

func NewUserDataTable(cfg *datatable.DataTaleConfig) *UserDataTable {
	dt := &UserDataTable{
		datatable.NewBaseDataTable(cfg, database.GetDatabase()),
	}
	dt.Query = dt.GetQuery()
	dt.Columns = dt.GetColumns()
	dt.ModifyDatatable()
	dt.ReducerFunc = dt.Reducer
	dt.BeforeProcessFunc = dt.BeforeProcess
	return dt
}

// GetQuery returns the main query for the datatable
// Override this method in child classes to define custom query
func (udt *UserDataTable) GetQuery() *gorm.DB {
	return udt.GetDB().Model(&model.User{})
}

// GetColumns returns the column definitions
// Override this method in child classes to define columns
func (udt *UserDataTable) GetColumns() []*datatable.ColumnDefinition {
	return []*datatable.ColumnDefinition{}
}

// ModifyDatatable is used to modify the datatable
func (udt *UserDataTable) ModifyDatatable() {

}

// Reducer is used to modify the datatable
func (udt *UserDataTable) Reducer(rows []interface{}) []interface{} {
	return rows
}

// BeforeProcess
func (udt *UserDataTable) BeforeProcess() {
	udt.Result = []model.User{}
}
