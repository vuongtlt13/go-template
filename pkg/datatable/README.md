# Datatable Package

Package datatable cung cấp chức năng xử lý datatable giống như Laravel DataTables cho Go với Fiber framework, được lấy cảm hứng từ implementation FastAPI mạnh mẽ.

## Tính năng

- ✅ Pagination (phân trang)
- ✅ Sorting (sắp xếp)
- ✅ Searching (tìm kiếm)
- ✅ Smart Search (tìm kiếm thông minh)
- ✅ Filtering (lọc dữ liệu)
- ✅ Custom Filter Rules (quy tắc lọc tùy chỉnh)
- ✅ Relations (preload relationships)
- ✅ Custom conditions (điều kiện tùy chỉnh)
- ✅ Additional Columns (cột bổ sung)
- ✅ Edit Columns (chỉnh sửa cột)
- ✅ Index Columns (cột index)
- ✅ Export functionality (chức năng xuất)
- ✅ Selected IDs filtering (lọc theo ID đã chọn)
- ✅ GORM integration
- ✅ Fiber integration
- ✅ **BaseDataTable inheritance pattern** (mới)

## Cài đặt

Package này đã được tích hợp sẵn trong project. Không cần cài đặt thêm.

## Cách sử dụng

### 1. Sử dụng BaseDataTable Pattern (Recommended)

Pattern này tương tự như FastAPI, cho phép tạo các DataTable con kế thừa từ BaseDataTable với các abstract methods.

#### Core Methods

BaseDataTable có 2 methods quan trọng nhất cần override:

- **`GetQuery()`**: Định nghĩa query chính cho datatable
- **`GetColumns()`**: Định nghĩa các columns và thuộc tính của chúng

#### Tạo UserDataTable

```go
// internal/datatable/user_datatable.go
package datatable

import (
    "yourproject/internal/schema"
    "gorm.io/gorm"
)

type UserDataTable struct {
    *BaseDataTable
}

func NewUserDataTable(db *gorm.DB) *UserDataTable {
    return &UserDataTable{
        BaseDataTable: NewBaseDataTable(db, &schema.UserInfo{}),
    }
}

// GetQuery returns the main query for users
func (udt *UserDataTable) GetQuery() *gorm.DB {
    return udt.db.Model(&schema.UserInfo{}).
        Where("deleted_at IS NULL").
        Order("id ASC")
}

// GetColumns returns the column definitions for users
func (udt *UserDataTable) GetColumns() []*ColumnDefinition {
    return []*ColumnDefinition{
        NewReadOnlyColumn("id", "ID"),
        NewSearchableColumn("email", "Email"),
        NewSearchableColumn("full_name", "Full Name"),
        NewOrderableColumn("is_active", "Status"),
        NewOrderableColumn("is_admin", "Role"),
        NewOrderableColumn("created_at", "Created At"),
        NewReadOnlyColumn("actions", "Actions"),
    }
}

// Helper functions for creating columns
func NewSearchableColumn(data, title string) *ColumnDefinition {
    return &ColumnDefinition{
        Data:       data,
        Title:      title,
        Searchable: true,
        Orderable:  true,
        Exportable: true,
        Printable:  true,
    }
}

func NewOrderableColumn(data, title string) *ColumnDefinition {
    return &ColumnDefinition{
        Data:       data,
        Title:      title,
        Searchable: false,
        Orderable:  true,
        Exportable: true,
        Printable:  true,
    }
}

func NewReadOnlyColumn(data, title string) *ColumnDefinition {
    return &ColumnDefinition{
        Data:       data,
        Title:      title,
        Searchable: false,
        Orderable:  false,
        Exportable: false,
        Printable:  false,
    }
}

// Optional: Override additional methods for custom logic
func (udt *UserDataTable) GetRelations() []string {
    return []string{"Roles", "Profile"}
}

func (udt *UserDataTable) GetAdditionalColumns() map[string]ProducerFunc {
    return map[string]ProducerFunc{
        "actions": func(record interface{}) interface{} {
            user := record.(*schema.UserInfo)
            return map[string]interface{}{
                "edit":   user.IsAdmin,
                "delete": true,
                "view":   true,
            }
        },
    }
}

func (udt *UserDataTable) GetEditColumns() map[string]ProducerFunc {
    return map[string]ProducerFunc{
        "is_active": func(record interface{}) interface{} {
            if record.(*schema.UserInfo).IsActive {
                return "Active"
            }
            return "Inactive"
        },
    }
}

func (udt *UserDataTable) ShouldIncludeIndex() bool {
    return true
}

func (udt *UserDataTable) GetMaxLimit() int {
    return 500
}
```

#### Sử dụng trong Handler

```go
// internal/handler/admin/user.go
func (h *UserHandler) GetUsers(c *fiber.Ctx) error {
    userDataTable := NewUserDataTable(h.db)
    return userDataTable.ProcessWithFiber(c)
}
```

#### Tạo RoleDataTable

```go
// internal/datatable/role_datatable.go
type RoleDataTable struct {
    *BaseDataTable
}

func NewRoleDataTable(db *gorm.DB) *RoleDataTable {
    return &RoleDataTable{
        BaseDataTable: NewBaseDataTable(db, &schema.RoleInfo{}),
    }
}

// GetQuery returns the main query for roles
func (rdt *RoleDataTable) GetQuery() *gorm.DB {
    return rdt.db.Model(&schema.RoleInfo{}).
        Where("deleted_at IS NULL").
        Order("name ASC")
}

// GetColumns returns the column definitions for roles
func (rdt *RoleDataTable) GetColumns() []*ColumnDefinition {
    return []*ColumnDefinition{
        NewReadOnlyColumn("id", "ID"),
        NewSearchableColumn("name", "Name"),
        NewSearchableColumn("description", "Description"),
        NewOrderableColumn("is_active", "Status"),
        NewOrderableColumn("is_system", "Type"),
        NewOrderableColumn("created_at", "Created At"),
        NewReadOnlyColumn("actions", "Actions"),
        NewReadOnlyColumn("permission_count", "Permissions"),
    }
}

func (rdt *RoleDataTable) GetRelations() []string {
    return []string{"Permissions"}
}

func (rdt *RoleDataTable) GetAdditionalColumns() map[string]ProducerFunc {
    return map[string]ProducerFunc{
        "actions": func(record interface{}) interface{} {
            role := record.(*schema.RoleInfo)
            return map[string]interface{}{
                "edit":   true,
                "delete": !role.IsSystem,
                "view":   true,
            }
        },
        "permission_count": func(record interface{}) interface{} {
            role := record.(*schema.RoleInfo)
            return len(role.Permissions)
        },
    }
}
```

#### So sánh với FastAPI

| FastAPI                         | Go BaseDataTable               |
| ------------------------------- | ------------------------------ |
| `query()`                       | `GetQuery()`                   |
| `get_columns()`                 | `GetColumns()`                 |
| `DataTableColumn`               | `ColumnDefinition`             |
| `select(User)`                  | `db.Model(&UserInfo{})`        |
| `filter(User.is_admin != True)` | `Where("is_admin != ?", true)` |

### 2. Cách đơn giản nhất (Legacy)

```go
// Trong handler
func (h *UserHandler) GetUsers(c *fiber.Ctx) error {
    return datatable.ProcessDatatable(c, h.db, &schema.UserInfo{}, func(dt *datatable.Datatable) {
        // Cấu hình searchable columns
        dt.AddSearchable("email")
        dt.AddSearchable("full_name")

        // Cấu hình orderable columns
        dt.AddOrderable("id")
        dt.AddOrderable("email")
        dt.AddOrderable("full_name")
        dt.AddOrderable("is_active")
        dt.AddOrderable("created_at")

        // Thêm relations nếu cần
        dt.AddRelation("Roles")

        // Thêm điều kiện tùy chỉnh
        dt.AddCondition("deleted_at IS NULL")
    })
}
```

### 2. Sử dụng với Options

```go
func (h *UserHandler) GetUsersWithOptions(c *fiber.Ctx) error {
    return datatable.ProcessDatatable(c, h.db, &schema.UserInfo{}, func(dt *datatable.Datatable) {
        // Cấu hình cơ bản
        dt.AddSearchable("email", "full_name")
        dt.AddOrderable("id", "email", "full_name", "created_at")
        dt.AddRelation("Roles")
        dt.AddCondition("deleted_at IS NULL")
    },
    datatable.WithMaxLimit(500),
    datatable.WithSmartSearch(true))
}
```

### 3. Sử dụng Advanced Features

```go
func (h *UserHandler) GetUsersAdvanced(c *fiber.Ctx) error {
    return datatable.ProcessDatatable(c, h.db, &schema.UserInfo{}, func(dt *datatable.Datatable) {
        // Cấu hình cơ bản
        dt.AddSearchable("email", "full_name")
        dt.AddOrderable("id", "email", "full_name", "created_at")
        dt.AddRelation("Roles")
        dt.AddCondition("deleted_at IS NULL")

        // Thêm cột bổ sung
        dt.AddAdditionalColumn("actions", func(record interface{}) interface{} {
            return map[string]interface{}{
                "edit":   true,
                "delete": true,
                "view":   true,
            }
        })

        // Chỉnh sửa cột hiện có
        dt.EditColumn("is_active", func(record interface{}) interface{} {
            if record.(*schema.UserInfo).IsActive {
                return "Active"
            }
            return "Inactive"
        })

        // Thêm cột index
        dt.AddIndexColumn("row_index")

        // Thêm custom filters
        dt.FilterColumn("status", func(value string) interface{} {
            switch value {
            case "active":
                return true
            case "inactive":
                return false
            default:
                return nil
            }
        })
    })
}
```

### 4. Sử dụng UseDatatable Decorator

```go
// Tạo decorator
var userDatatable = datatable.NewUseDatatable(500, true)

// Trong handler
func (h *UserHandler) GetUsersWithDecorator(c *fiber.Ctx) error {
    return userDatatable.Call(c, h.db, &schema.UserInfo{}, func(dt *datatable.Datatable) {
        dt.AddSearchable("email", "full_name")
        dt.AddOrderable("id", "email", "full_name", "created_at")
        dt.AddRelation("Roles")
        dt.AddCondition("deleted_at IS NULL")
    })
}
```

### 5. Cách chi tiết hơn

```go
// Trong handler
func (h *UserHandler) GetUsers(c *fiber.Ctx) error {
    // Parse request
    req, err := datatable.ParseRequest(c)
    if err != nil {
        return datatable.SendErrorResponse(c, 0, "Invalid request format")
    }

    // Tạo datatable với options
    dt := datatable.New(h.db, &schema.UserInfo{},
        datatable.WithMaxLimit(500),
        datatable.WithSmartSearch(true),
    )
    dt.SetRequest(req)

    // Cấu hình
    dt.AddSearchable("email", "full_name")
    dt.AddOrderable("id", "email", "full_name", "created_at")
    dt.AddRelation("Roles")
    dt.AddCondition("deleted_at IS NULL")

    // Xử lý
    response, err := dt.Process()
    if err != nil {
        return datatable.SendErrorResponse(c, req.Draw, err.Error())
    }

    // Trả về response
    return datatable.SendResponse(c, response)
}
```

## Request Format

Datatable nhận request theo format của DataTables với các extension:

### JSON Request

```json
{
  "draw": 1,
  "start": 0,
  "length": 10,
  "search": {
    "value": "john",
    "regex": false
  },
  "order": [
    {
      "column": 0,
      "dir": "asc"
    }
  ],
  "columns": [
    {
      "data": "id",
      "name": "id",
      "searchable": true,
      "orderable": true,
      "exportable": true,
      "printable": true,
      "class_name": "text-center",
      "search": {
        "value": "",
        "regex": false
      }
    }
  ],
  "action": "ajax",
  "keyword": "john doe",
  "skip": 0,
  "limit": 25,
  "sort_bys": ["email", "created_at"],
  "sort_types": ["asc", "desc"],
  "selected_ids": [1, 2, 3]
}
```

### Query Parameters

```
?q=john%20doe&s=0&ipp=25&action=ajax&sb[]=email&sd[]=asc&ids=[1,2,3]
```

## Response Format

```json
{
    "draw": 1,
    "recordsTotal": 100,
    "recordsFiltered": 25,
    "data": [
        {
            "id": 1,
            "email": "john@example.com",
            "full_name": "John Doe",
            "is_active": "Active",
            "created_at": "2024-01-01T00:00:00Z",
            "row_index": 1,
            "actions": {
                "edit": true,
                "delete": true,
                "view": true
            }
        }
    ],
    "items": [...],
    "others": {
        "total_pages": 4,
        "current_page": 1
    }
}
```

## API Reference

### Options

#### WithMaxLimit(maxLimit int) Option

Set giới hạn tối đa số records trả về.

#### WithSmartSearch(smartSearch bool) Option

Bật/tắt smart search (tìm kiếm theo từng từ riêng biệt).

### Datatable Methods

#### New(db *gorm.DB, model interface{}, options ...Option) *Datatable

Tạo instance datatable mới với options.

#### SetRequest(req *Request) *Datatable

Set request data.

#### AddSearchable(column string) \*Datatable

Thêm column có thể search.

#### AddOrderable(column string) \*Datatable

Thêm column có thể sort.

#### AddRelation(relation string) \*Datatable

Thêm relation để preload.

#### AddCondition(condition string, args ...interface{}) \*Datatable

Thêm điều kiện WHERE.

#### AddAdditionalColumn(columnName string, producer ProducerFunc) \*Datatable

Thêm cột bổ sung với function tạo dữ liệu.

#### EditColumn(columnName string, producer ProducerFunc) \*Datatable

Chỉnh sửa cột hiện có.

#### AddIndexColumn(columnName string) \*Datatable

Thêm cột index.

#### FilterColumn(columnName string, filter FilterFunc) \*Datatable

Thêm custom filter cho column.

#### Process() (\*Response, error)

Xử lý datatable và trả về response.

### Helper Functions

#### ParseRequest(c *fiber.Ctx) (*Request, error)

Parse request từ Fiber context.

#### ProcessDatatable(c *fiber.Ctx, db *gorm.DB, model interface{}, configurator func(\*Datatable), options ...Option) error

Xử lý datatable với cấu hình callback và options.

#### SendResponse(c *fiber.Ctx, response *Response) error

Gửi response.

#### SendErrorResponse(c \*fiber.Ctx, draw int, errorMsg string) error

Gửi error response.

#### NewUseDatatable(maxLimit int, smartSearch bool, options ...Option) \*UseDatatable

Tạo UseDatatable decorator.

## Ví dụ thực tế

### User Management với Advanced Features

```go
// internal/handler/admin/user.go
func (h *UserHandler) GetUsersAdvanced(c *fiber.Ctx) error {
    return datatable.ProcessDatatable(c, h.db, &schema.UserInfo{}, func(dt *datatable.Datatable) {
        // Searchable columns
        dt.AddSearchable("email")
        dt.AddSearchable("full_name")

        // Orderable columns
        dt.AddOrderable("id")
        dt.AddOrderable("email")
        dt.AddOrderable("full_name")
        dt.AddOrderable("is_active")
        dt.AddOrderable("is_admin")
        dt.AddOrderable("created_at")

        // Relations
        dt.AddRelation("Roles")

        // Conditions
        dt.AddCondition("deleted_at IS NULL")

        // Additional columns
        dt.AddAdditionalColumn("actions", func(record interface{}) interface{} {
            user := record.(*schema.UserInfo)
            return map[string]interface{}{
                "edit":   user.IsAdmin,
                "delete": true,
                "view":   true,
            }
        })

        // Edit columns
        dt.EditColumn("is_active", func(record interface{}) interface{} {
            if record.(*schema.UserInfo).IsActive {
                return "Active"
            }
            return "Inactive"
        })

        // Index column
        dt.AddIndexColumn("row_index")

        // Custom filters
        dt.FilterColumn("status", func(value string) interface{} {
            switch value {
            case "active":
                return true
            case "inactive":
                return false
            default:
                return nil
            }
        })
    },
    datatable.WithMaxLimit(500),
    datatable.WithSmartSearch(true))
}
```

### Product Management với Export

```go
// internal/handler/admin/product.go
func (h *ProductHandler) GetProducts(c *fiber.Ctx) error {
    return datatable.ProcessDatatable(c, h.db, &schema.ProductInfo{}, func(dt *datatable.Datatable) {
        // Searchable columns
        dt.AddSearchable("name")
        dt.AddSearchable("sku")
        dt.AddSearchable("description")

        // Orderable columns
        dt.AddOrderable("id")
        dt.AddOrderable("name")
        dt.AddOrderable("price")
        dt.AddOrderable("stock")
        dt.AddOrderable("created_at")

        // Relations
        dt.AddRelation("Category")
        dt.AddRelation("Brand")

        // Conditions
        dt.AddCondition("is_active = ?", true)

        // Additional columns for export
        dt.AddAdditionalColumn("export_price", func(record interface{}) interface{} {
            product := record.(*schema.ProductInfo)
            return fmt.Sprintf("$%.2f", product.Price)
        })
    })
}
```

## Best Practices

1. **Sử dụng ProcessDatatable helper**: Đây là cách đơn giản và an toàn nhất.

2. **Cấu hình searchable/orderable columns**: Chỉ cho phép search/sort những column cần thiết.

3. **Sử dụng relations**: Preload relations để tránh N+1 query.

4. **Thêm conditions**: Luôn thêm điều kiện cơ bản như `deleted_at IS NULL`.

5. **Error handling**: Luôn xử lý lỗi và trả về error response phù hợp.

6. **Performance**: Sử dụng index cho các column thường xuyên search/sort.

7. **Smart Search**: Sử dụng smart search để tìm kiếm hiệu quả hơn.

8. **Custom Filters**: Sử dụng custom filters cho logic phức tạp.

9. **Additional Columns**: Sử dụng additional columns cho dữ liệu tính toán.

10. **UseDatatable Decorator**: Sử dụng decorator cho code tái sử dụng.

## Troubleshooting

### Lỗi thường gặp

1. **Column not found**: Đảm bảo tên column đúng với database schema.
2. **Relation not found**: Kiểm tra tên relation trong model.
3. **Invalid condition**: Kiểm tra syntax của điều kiện WHERE.
4. **Filter function error**: Kiểm tra logic trong FilterFunc.

### Debug

```go
// Lấy query để debug
dt := datatable.New(db, model)
query := dt.GetQuery()
fmt.Println(query.ToSQL(func(tx *gorm.DB) *gorm.DB {
    return tx.Find(&[]YourModel{})
}))
```

## Migration từ Laravel DataTables

Nếu bạn đang sử dụng Laravel DataTables, việc migrate sang Go sẽ khá đơn giản:

### Laravel

```php
return DataTables::of(User::query())
    ->addColumn('actions', function($user) {
        return view('users.actions', compact('user'));
    })
    ->filterColumn('email', function($query, $keyword) {
        $query->where('email', 'like', "%{$keyword}%");
    })
    ->make(true);
```

### Go

```go
return datatable.ProcessDatatable(c, db, &schema.UserInfo{}, func(dt *datatable.Datatable) {
    dt.AddSearchable("email")
    dt.AddOrderable("id", "email", "created_at")
    dt.AddAdditionalColumn("actions", func(record interface{}) interface{} {
        return map[string]interface{}{"edit": true, "delete": true}
    })
    dt.FilterColumn("email", func(value string) interface{} {
        return "%" + value + "%"
    })
})
```

## So sánh với FastAPI Implementation

Package này được lấy cảm hứng từ implementation FastAPI mạnh mẽ và cung cấp các tính năng tương tự:

| Tính năng            | FastAPI | Go  |
| -------------------- | ------- | --- |
| Smart Search         | ✅      | ✅  |
| Custom Filters       | ✅      | ✅  |
| Additional Columns   | ✅      | ✅  |
| Edit Columns         | ✅      | ✅  |
| Index Columns        | ✅      | ✅  |
| Export Functionality | ✅      | ✅  |
| Selected IDs Filter  | ✅      | ✅  |
| Decorator Pattern    | ✅      | ✅  |
| Options Pattern      | ✅      | ✅  |
