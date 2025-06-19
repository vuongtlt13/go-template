package datatable

// DataTableColumn represents a column definition
type DataTableColumn struct {
	Data        string
	Searchable  bool
	Orderable   bool
	Exportable  bool
	Printable   bool
	ColumnAlias string
	Filter      FilterFunc
}

func NewDataTableColumn(data string, searchable bool, orderable bool, exportable bool, printable bool, columnAlias string) *DataTableColumn {

	return &DataTableColumn{
		Data:        data,
		Searchable:  searchable,
		Orderable:   orderable,
		Exportable:  exportable,
		Printable:   printable,
		ColumnAlias: columnAlias,
		Filter:      nil,
	}
}
