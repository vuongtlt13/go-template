package datatable

import (
	"yourapp/internal/model"
	"yourapp/pkg/database"
	"yourapp/pkg/datatable"

	"gorm.io/gorm"
)

type UserDataTable struct {
	*datatable.BaseDataTable
}

func NewUserDataTable(cfg *datatable.DataTaleConfig) *UserDataTable {
	dt := &UserDataTable{
		BaseDataTable: datatable.NewBaseDataTable(
			cfg,
			database.GetDatabase(),
		),
	}
	dt.SetOverride(dt)
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
	return []*datatable.ColumnDefinition{
		{Data: "email", Searchable: true, Orderable: true, Exportable: true, Printable: true},
		{Data: "full_name", Searchable: true, Orderable: true, Exportable: true, Printable: true},
		{Data: "is_active", Searchable: false, Orderable: true, Exportable: true, Printable: true},
		{Data: "is_admin", Searchable: false, Orderable: true, Exportable: true, Printable: true},
	}
}

// ModifyDatatable is used to modify the datatable
func (udt *UserDataTable) ModifyDatatable() {}

// AfterProcess
func (udt *UserDataTable) AfterProcess() {}

// BeforeProcess
func (udt *UserDataTable) BeforeProcess() {
	udt.Result = []model.User{}
}
