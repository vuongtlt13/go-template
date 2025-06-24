# Go CRUD Code Generator

## Tổng quan

Generator này giúp tự động sinh code cho các module CRUD (repository, schema, datatable, service, handler), chuẩn multi-layer Go, từ một Go model duy nhất. Generator hỗ trợ chuẩn hóa kiến trúc, tăng tốc phát triển, giảm lặp lại code và đảm bảo đồng bộ giữa các layer.

## Cách hoạt động

1. **Phân tích model Go** (struct trong `internal/model/<entity>.go`) bằng AST, lấy thông tin field, tag, validate, relation, ...
2. **Sinh code cho từng component** (repository, schema, datatable, service, handler) bằng cách render Go template với dữ liệu entity đã parse.
3. **Tự động format code** với `goimports` sau khi generate.
4. **Có thể in ra màn hình (`--print`) hoặc ghi file đúng thư mục layer.**

## Cấu trúc thư mục

```
pkg/generator/
  core/                # Logic parse model, base generator, entity, ...
  backend/             # Generator cho các component backend
    repository_generator.go
    schema_generator.go
    datatable_generator.go
    service_generator.go
    handler_generator.go
  templates/
    backend/           # Template cho từng component backend
      repository.go.tmpl
      schema.go.tmpl
      datatable.go.tmpl
      service.go.tmpl
      handler.go.tmpl
  main.go             # CLI entry point
```

## Các component hỗ trợ

- **repository**: Truy vấn DB, CRUD cho entity
- **schema**: Định nghĩa request/response struct cho API
- **datatable**: Listing, paging, filter cho entity
- **service**: Logic nghiệp vụ, validate, xử lý phức tạp
- **handler**: Xử lý HTTP API với Fiber framework

## Vị trí file mẫu (template)

- Template cho từng component nằm ở: `pkg/generator/templates/backend/<component>.go.tmpl`
- Có thể sửa template để customize code sinh ra.

## Quy trình generate từ model

1. **Viết model Go** trong `internal/model/<entity>.go` (ví dụ: `User`)
2. **Chạy generator**:
   - Sinh repository: `go run pkg/generator/main.go --cmd=generate --model=User --component=repository`
   - Sinh schema: `go run pkg/generator/main.go --cmd=generate --model=User --component=schema`
   - Sinh datatable: `go run pkg/generator/main.go --cmd=generate --model=User --component=datatable`
   - Sinh service: `go run pkg/generator/main.go --cmd=generate --model=User --component=service`
   - Sinh handler: `go run pkg/generator/main.go --cmd=generate --model=User --component=handler`
   - Thêm `--print` để chỉ in ra màn hình, không ghi file.
3. **File sinh ra sẽ nằm ở đúng thư mục layer:**
   - `internal/repository/user.go`
   - `internal/schema/user.go`
   - `internal/datatable/user.go`
   - `internal/service/user.go`
   - `internal/handler/admin/user.go`
4. **Code sinh ra đã được format chuẩn Go.**

## Logic mapping

- Generator tự động lấy tên module từ `go.mod` để import đúng path.
- Chỉ sinh cột datatable cho field primitive, không sinh cho relation/json:"-".
- Schema sinh theo mẫu chuẩn Go, có thể mở rộng filter nếu cần.
- Service layer chứa logic nghiệp vụ, validate, xử lý phức tạp.
- Handler layer xử lý HTTP request/response với Fiber framework và Swagger docs.
- Có thể mở rộng thêm component (router, api, register routes, ...) bằng cách thêm template và generator tương ứng.

## Quy trình generate đầy đủ các layer cho 1 model (ví dụ: user)

Khi muốn tạo đầy đủ các layer cho 1 model (ví dụ: user), bạn cần thực hiện các bước sau:

1. **Thêm file define router và register tại thư mục `internal/api/admin`.**

   - File: `internal/api/admin/user.go`
   - Mục đích: Định nghĩa các route, group, middleware cho API user.

2. **Implement handler cho các API tại thư mục `internal/handler/admin`.**

   - File: `internal/handler/admin/user.go` ✅ **Đã có generator**
   - Mục đích: Xử lý logic nhận request, trả response cho từng API (CRUD, datatable, ...).

3. **Implement datatable để lấy danh sách record có paging tại thư mục `internal/datatable`.**

   - File: `internal/datatable/user.go` ✅ **Đã có generator**
   - Mục đích: Định nghĩa logic truy vấn, filter, paging cho listing user.

4. **Implement repository tại thư mục `internal/repository`.**

   - File: `internal/repository/user.go` ✅ **Đã có generator**
   - Mục đích: Truy vấn DB, CRUD cho entity user.

5. **Khai báo schema tại thư mục `internal/schema`.**

   - File: `internal/schema/user.go` ✅ **Đã có generator**
   - Mục đích: Định nghĩa struct request/response cho API user.

6. **Implement service xử lý logic tại thư mục `internal/service`.**

   - File: `internal/service/user.go` ✅ **Đã có generator**
   - Mục đích: Chứa logic nghiệp vụ, validate, xử lý phức tạp cho user.

7. **Thêm bản dịch tiếng Anh của model tại `i18n/locales/en/models`.**

   - File: `i18n/locales/en/models/user.json`
   - Mục đích: Định nghĩa các key/value bản dịch cho user (label, message, ...).

8. **Đăng ký các routes của User tại file `internal/routes/admin.go`.**
   - Hàm: `Register`
   - Mục đích: Đăng ký route user vào router tổng của hệ thống.

> **Lưu ý:** Các bước 2-6 đã có generator tự động. Chỉ cần chạy lệnh generate cho từng component.

## Mở rộng

- Để thêm component mới: tạo template và generator, mapping vào CLI.
- Có thể customize template cho từng project.
- Có thể tích hợp thêm validate, i18n, test, ...

## Đóng góp

- PR, issue, góp ý đều welcome!

---

**Ví dụ sử dụng:**

```
go run pkg/generator/main.go --cmd=generate --model=User --component=repository
```

**Sinh nhiều component:** (tương lai)

```
go run pkg/generator/main.go --cmd=generate --model=User --component=repository,schema,datatable,service,handler
```

**In ra màn hình:**

```
go run pkg/generator/main.go --cmd=generate --model=User --component=schema --print
```
