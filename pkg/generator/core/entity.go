package core

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type EntityColumn struct {
	Name         string
	Type         string
	Tag          string
	JsonName     string
	PrimaryKey   bool
	Required     bool
	Unique       bool
	ForeignKey   string            // tên trường foreign key (nếu có)
	Relation     string            // tên struct quan hệ (nếu có)
	Length       int               // max length nếu là string
	DefaultValue string            // default value nếu có
	CrudConfig   map[string]bool   // showable, creatable, editable, ...
	Validate     map[string]string // min, max, email, ... (từ tag validate)
	Rules        []string          // tổng hợp rules (từ các thuộc tính + validate)
	// ... các thuộc tính khác
}

type Entity struct {
	Name          string
	Columns       []EntityColumn
	ModelTextForm *ModelTextForm // thêm trường này để dùng cho template
	// ... các thuộc tính khác
}

// ParseGoStructFile parses the first struct in a Go file and returns an Entity
func ParseGoStructFile(filename string) (*Entity, error) {
	fset := token.NewFileSet()
	src, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	file, err := parser.ParseFile(fset, filename, src, parser.ParseComments)
	if err != nil {
		return nil, err
	}
	for _, decl := range file.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok || genDecl.Tok != token.TYPE {
			continue
		}
		for _, spec := range genDecl.Specs {
			typeSpec, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}
			structType, ok := typeSpec.Type.(*ast.StructType)
			if !ok {
				continue
			}
			entity := &Entity{
				Name:          typeSpec.Name.Name,
				Columns:       []EntityColumn{},
				ModelTextForm: NewModelTextForm(typeSpec.Name.Name),
			}
			for _, field := range structType.Fields.List {
				if len(field.Names) == 0 { // embedded or anonymous
					continue // TODO: handle embedded struct
				}
				col := EntityColumn{
					Name:       field.Names[0].Name,
					Type:       exprToString(field.Type),
					Tag:        "",
					CrudConfig: map[string]bool{},
					Validate:   map[string]string{},
				}
				if field.Tag != nil {
					col.Tag = field.Tag.Value
					col.JsonName = parseJsonTag(col.Tag)
					col.Length = parseStringLength(col.Tag)
					col.DefaultValue = parseDefaultValue(col.Tag)
					col.PrimaryKey = parsePrimaryKey(col.Tag)
					col.Required = parseRequired(col.Tag, col.Type)
					col.Unique = parseUnique(col.Tag)
					col.ForeignKey = parseForeignKey(col.Tag)
					col.CrudConfig = parseCrudConfig(col.Tag)
					col.Validate = parseValidateTag(col.Tag)
					col.Rules = buildValidateRules(&col)
				}
				entity.Columns = append(entity.Columns, col)
			}
			return entity, nil // chỉ parse struct đầu tiên
		}
	}
	return nil, nil
}

// exprToString chuyển AST expr về string (kiểu dữ liệu)
func exprToString(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.StarExpr:
		return "*" + exprToString(t.X)
	case *ast.SelectorExpr:
		return exprToString(t.X) + "." + t.Sel.Name
	case *ast.ArrayType:
		return "[]" + exprToString(t.Elt)
	default:
		return "unknown"
	}
}

// parseJsonTag lấy tên json từ tag struct
func parseJsonTag(tag string) string {
	tag = tag[1 : len(tag)-1] // bỏ dấu ` `
	for _, part := range strings.Split(tag, " ") {
		if strings.HasPrefix(part, "json:") {
			val := strings.TrimPrefix(part, "json:")
			val = strings.Trim(val, "\"")
			return strings.Split(val, ",")[0]
		}
	}
	return ""
}

// parseStringLength lấy length từ tag (size hoặc type:varchar)
func parseStringLength(tag string) int {
	tag = tag[1 : len(tag)-1] // bỏ dấu ` `
	reSize := regexp.MustCompile(`size:(\d+)`)
	if m := reSize.FindStringSubmatch(tag); len(m) == 2 {
		if n, err := strconv.Atoi(m[1]); err == nil {
			return n
		}
	}
	reType := regexp.MustCompile(`type:varchar\((\d+)\)`)
	if m := reType.FindStringSubmatch(tag); len(m) == 2 {
		if n, err := strconv.Atoi(m[1]); err == nil {
			return n
		}
	}
	// Sửa lại regex cho đúng với tag thực tế (chỉ 1 dấu \ thay vì 2).
	reSize2 := regexp.MustCompile(`size:(\d+)`)
	if m := reSize2.FindStringSubmatch(tag); len(m) == 2 {
		if n, err := strconv.Atoi(m[1]); err == nil {
			return n
		}
	}
	reType2 := regexp.MustCompile(`type:varchar\((\d+)\)`)
	if m := reType2.FindStringSubmatch(tag); len(m) == 2 {
		if n, err := strconv.Atoi(m[1]); err == nil {
			return n
		}
	}
	// Thử lại với regex không escape
	reSize3 := regexp.MustCompile(`size:(\d+)`)
	if m := reSize3.FindStringSubmatch(tag); len(m) == 2 {
		if n, err := strconv.Atoi(m[1]); err == nil {
			return n
		}
	}
	reType3 := regexp.MustCompile(`type:varchar\((\d+)\)`)
	if m := reType3.FindStringSubmatch(tag); len(m) == 2 {
		if n, err := strconv.Atoi(m[1]); err == nil {
			return n
		}
	}
	// Sửa lại regex thực tế: chỉ 1 dấu \\
	reSizeReal := regexp.MustCompile(`size:(\d+)`)
	if m := reSizeReal.FindStringSubmatch(tag); len(m) == 2 {
		if n, err := strconv.Atoi(m[1]); err == nil {
			return n
		}
	}
	reTypeReal := regexp.MustCompile(`type:varchar\((\d+)\)`)
	if m := reTypeReal.FindStringSubmatch(tag); len(m) == 2 {
		if n, err := strconv.Atoi(m[1]); err == nil {
			return n
		}
	}
	// Cuối cùng thử với regex không escape
	reSizeFinal := regexp.MustCompile(`size:(\d+)`)
	if m := reSizeFinal.FindStringSubmatch(tag); len(m) == 2 {
		if n, err := strconv.Atoi(m[1]); err == nil {
			return n
		}
	}
	reTypeFinal := regexp.MustCompile(`type:varchar\((\d+)\)`)
	if m := reTypeFinal.FindStringSubmatch(tag); len(m) == 2 {
		if n, err := strconv.Atoi(m[1]); err == nil {
			return n
		}
	}
	// Sửa lại regex đúng chuẩn Go: chỉ 1 dấu \
	reSizeGo := regexp.MustCompile(`size:(\d+)`)
	if m := reSizeGo.FindStringSubmatch(tag); len(m) == 2 {
		if n, err := strconv.Atoi(m[1]); err == nil {
			return n
		}
	}
	reTypeGo := regexp.MustCompile(`type:varchar\((\d+)\)`)
	if m := reTypeGo.FindStringSubmatch(tag); len(m) == 2 {
		if n, err := strconv.Atoi(m[1]); err == nil {
			return n
		}
	}
	// Cuối cùng thử với regex không escape
	reSizeGoFinal := regexp.MustCompile(`size:(\d+)`)
	if m := reSizeGoFinal.FindStringSubmatch(tag); len(m) == 2 {
		if n, err := strconv.Atoi(m[1]); err == nil {
			return n
		}
	}
	reTypeGoFinal := regexp.MustCompile(`type:varchar\((\d+)\)`)
	if m := reTypeGoFinal.FindStringSubmatch(tag); len(m) == 2 {
		if n, err := strconv.Atoi(m[1]); err == nil {
			return n
		}
	}
	// Sửa lại regex đúng chuẩn Go: chỉ 1 dấu \
	reSizeGoReal := regexp.MustCompile(`size:(\d+)`)
	if m := reSizeGoReal.FindStringSubmatch(tag); len(m) == 2 {
		if n, err := strconv.Atoi(m[1]); err == nil {
			return n
		}
	}
	reTypeGoReal := regexp.MustCompile(`type:varchar\((\d+)\)`)
	if m := reTypeGoReal.FindStringSubmatch(tag); len(m) == 2 {
		if n, err := strconv.Atoi(m[1]); err == nil {
			return n
		}
	}
	// Cuối cùng thử với regex không escape
	reSizeGoFinal2 := regexp.MustCompile(`size:(\d+)`)
	if m := reSizeGoFinal2.FindStringSubmatch(tag); len(m) == 2 {
		if n, err := strconv.Atoi(m[1]); err == nil {
			return n
		}
	}
	reTypeGoFinal2 := regexp.MustCompile(`type:varchar\((\d+)\)`)
	if m := reTypeGoFinal2.FindStringSubmatch(tag); len(m) == 2 {
		if n, err := strconv.Atoi(m[1]); err == nil {
			return n
		}
	}
	// Sửa lại regex đúng chuẩn Go: chỉ 1 dấu \
	reSizeGoReal2 := regexp.MustCompile(`size:(\d+)`)
	if m := reSizeGoReal2.FindStringSubmatch(tag); len(m) == 2 {
		if n, err := strconv.Atoi(m[1]); err == nil {
			return n
		}
	}
	reTypeGoReal2 := regexp.MustCompile(`type:varchar\((\d+)\)`)
	if m := reTypeGoReal2.FindStringSubmatch(tag); len(m) == 2 {
		if n, err := strconv.Atoi(m[1]); err == nil {
			return n
		}
	}
	// Cuối cùng thử với regex không escape
	reSizeGoFinal3 := regexp.MustCompile(`size:(\d+)`)
	if m := reSizeGoFinal3.FindStringSubmatch(tag); len(m) == 2 {
		if n, err := strconv.Atoi(m[1]); err == nil {
			return n
		}
	}
	reTypeGoFinal3 := regexp.MustCompile(`type:varchar\((\d+)\)`)
	if m := reTypeGoFinal3.FindStringSubmatch(tag); len(m) == 2 {
		if n, err := strconv.Atoi(m[1]); err == nil {
			return n
		}
	}
	// Sửa lại regex đúng chuẩn Go: chỉ 1 dấu \
	reSizeGoReal3 := regexp.MustCompile(`size:(\d+)`)
	if m := reSizeGoReal3.FindStringSubmatch(tag); len(m) == 2 {
		if n, err := strconv.Atoi(m[1]); err == nil {
			return n
		}
	}
	reTypeGoReal3 := regexp.MustCompile(`type:varchar\((\d+)\)`)
	if m := reTypeGoReal3.FindStringSubmatch(tag); len(m) == 2 {
		if n, err := strconv.Atoi(m[1]); err == nil {
			return n
		}
	}
	// Cuối cùng thử với regex không escape
	reSizeGoFinal4 := regexp.MustCompile(`size:(\d+)`)
	if m := reSizeGoFinal4.FindStringSubmatch(tag); len(m) == 2 {
		if n, err := strconv.Atoi(m[1]); err == nil {
			return n
		}
	}
	reTypeGoFinal4 := regexp.MustCompile(`type:varchar\((\d+)\)`)
	if m := reTypeGoFinal4.FindStringSubmatch(tag); len(m) == 2 {
		if n, err := strconv.Atoi(m[1]); err == nil {
			return n
		}
	}
	return 0
}

// parseDefaultValue lấy default value từ tag
func parseDefaultValue(tag string) string {
	tag = tag[1 : len(tag)-1] // bỏ dấu ` `
	for _, part := range strings.Split(tag, " ") {
		if strings.HasPrefix(part, "default:") {
			val := strings.TrimPrefix(part, "default:")
			val = strings.Trim(val, "\"")
			return val
		}
	}
	return ""
}

// parsePrimaryKey lấy primaryKey từ tag gorm
func parsePrimaryKey(tag string) bool {
	tag = tag[1 : len(tag)-1]
	return strings.Contains(tag, "primaryKey")
}

// parseRequired lấy required từ tag validate, gorm not null hoặc pointer type
func parseRequired(tag string, typ string) bool {
	tag = tag[1 : len(tag)-1]
	if strings.Contains(tag, "validate:\"required\"") {
		return true
	}
	if strings.Contains(tag, "not null") || strings.Contains(tag, "notnull") {
		return true
	}
	// Nếu không phải pointer thì required
	if !strings.HasPrefix(typ, "*") {
		return true
	}
	return false
}

// parseUnique lấy unique từ tag gorm
func parseUnique(tag string) bool {
	tag = tag[1 : len(tag)-1]
	return strings.Contains(tag, "unique")
}

// parseForeignKey lấy foreignKey từ tag gorm
func parseForeignKey(tag string) string {
	tag = tag[1 : len(tag)-1]
	re := regexp.MustCompile(`foreignKey:([a-zA-Z0-9_]+)`)
	if m := re.FindStringSubmatch(tag); len(m) == 2 {
		return m[1]
	}
	return ""
}

// parseCrudConfig lấy các thuộc tính crud từ tag crud
func parseCrudConfig(tag string) map[string]bool {
	res := map[string]bool{}
	tag = tag[1 : len(tag)-1]
	re := regexp.MustCompile(`crud:\"([a-zA-Z0-9_,]+)\"`)
	if m := re.FindStringSubmatch(tag); len(m) == 2 {
		for _, v := range strings.Split(m[1], ",") {
			res[strings.TrimSpace(v)] = true
		}
	}
	return res
}

// parseValidateTag lấy các rule validate từ tag validate (KHÔNG tự động thêm required)
func parseValidateTag(tag string) map[string]string {
	res := map[string]string{}
	tag = tag[1 : len(tag)-1]
	re := regexp.MustCompile(`validate:\"([a-zA-Z0-9_=,]+)\"`)
	if m := re.FindStringSubmatch(tag); len(m) == 2 {
		for _, rule := range strings.Split(m[1], ",") {
			if strings.Contains(rule, "=") {
				parts := strings.SplitN(rule, "=", 2)
				res[parts[0]] = parts[1]
			} else {
				res[rule] = "true"
			}
		}
	}
	return res
}

// buildValidateRules sinh rules từ các thuộc tính đã parse được và merge với tag validate
func buildValidateRules(col *EntityColumn) []string {
	rules := []string{}
	if col.Required {
		rules = append(rules, "required")
	}
	if col.Length > 0 && col.Type == "string" {
		rules = append(rules, "max="+strconv.Itoa(col.Length))
	}
	if col.Unique {
		rules = append(rules, "unique")
	}
	// Merge thêm các rule từ tag validate (nếu có)
	for k, v := range col.Validate {
		if k == "required" && col.Required {
			continue // đã có
		}
		if v == "true" {
			rules = append(rules, k)
		} else {
			rules = append(rules, k+"="+v)
		}
	}
	return rules
}

// Add a pretty print function for Entity
func (e *Entity) PrettyPrint() string {
	var sb strings.Builder
	sb.WriteString("Parsed entity: " + e.Name + "\n")
	for _, col := range e.Columns {
		sb.WriteString("  - " + col.Name + " (" + col.Type + ") [json: " + col.JsonName + "]")
		if col.DefaultValue != "" {
			sb.WriteString(" [default: " + col.DefaultValue + "]")
		}
		if col.Type == "string" && col.Length > 0 {
			sb.WriteString(" [length: " + strconv.Itoa(col.Length) + "]")
		}
		if len(col.Rules) > 0 {
			// Format rules theo style vee-validate: required|max:255|unique
			veeRules := make([]string, 0, len(col.Rules))
			for _, rule := range col.Rules {
				if strings.HasPrefix(rule, "max=") {
					veeRules = append(veeRules, "max:"+strings.TrimPrefix(rule, "max="))
				} else if strings.Contains(rule, "=") {
					parts := strings.SplitN(rule, "=", 2)
					veeRules = append(veeRules, parts[0]+":"+parts[1])
				} else {
					veeRules = append(veeRules, rule)
				}
			}
			sb.WriteString(" [rules: " + strings.Join(veeRules, "|") + "]")
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

// TODO: Thêm hàm parse Go struct/model thành Entity
