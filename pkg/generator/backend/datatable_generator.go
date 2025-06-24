package backend

import (
	"path/filepath"
	"strings"
	"yourapp/pkg/generator/core"
)

type DataTableGenerator struct {
	*core.BaseGenerator
}

func NewDataTableGenerator(entity *core.Entity, templatePath string, force bool) core.Generator {
	return &DataTableGenerator{
		BaseGenerator: &core.BaseGenerator{
			Entity:       entity,
			TemplatePath: templatePath,
			Force:        force,
		},
	}
}

func (g *DataTableGenerator) TemplateFile() string { return "backend/datatable.go.tmpl" }
func (g *DataTableGenerator) OutputFilename() string {
	return g.Entity.ModelTextForm.ModelNameSnakeCase + ".go"
}
func (g *DataTableGenerator) ComponentFolder() string { return "internal/datatable" }
func (g *DataTableGenerator) InputData() map[string]interface{} {
	moduleName, _ := core.GetGoModuleName()
	columns := []core.EntityColumn{}
	for _, col := range g.Entity.Columns {
		if col.JsonName == "-" || col.JsonName == "" {
			continue
		}
		// Loại bỏ relation (slice hoặc struct không phải time.Time)
		if strings.HasPrefix(col.Type, "[]") && col.Type != "[]byte" {
			continue
		}
		if col.Type != "string" && col.Type != "int" && col.Type != "bool" && col.Type != "float64" && col.Type != "float32" && col.Type != "uint" && col.Type != "uint64" && col.Type != "int64" && col.Type != "int32" && col.Type != "time.Time" {
			continue
		}
		columns = append(columns, col)
	}
	return map[string]interface{}{
		"Entity":     g.Entity,
		"ModuleName": moduleName,
		"Columns":    columns,
	}
}

func (g *DataTableGenerator) Generate() (string, error) {
	return RenderTemplateWithFuncMap(
		filepath.Join(g.TemplatePath, g.TemplateFile()),
		g.InputData(),
		nil, // Không cần custom funcMap cho datatable
	)
}
