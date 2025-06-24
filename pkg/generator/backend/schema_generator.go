package backend

import (
	"path/filepath"
	"strings"
	"text/template"
	"yourapp/pkg/generator/core"
)

type SchemaGenerator struct {
	*core.BaseGenerator
}

func NewSchemaGenerator(entity *core.Entity, templatePath string, force bool) core.Generator {
	return &SchemaGenerator{
		BaseGenerator: &core.BaseGenerator{
			Entity:       entity,
			TemplatePath: templatePath,
			Force:        force,
		},
	}
}

func (g *SchemaGenerator) TemplateFile() string { return "backend/schema.go.tmpl" }
func (g *SchemaGenerator) OutputFilename() string {
	return g.Entity.ModelTextForm.ModelNameSnakeCase + ".go"
}
func (g *SchemaGenerator) ComponentFolder() string { return "internal/schema" }
func (g *SchemaGenerator) InputData() map[string]interface{} {
	moduleName, _ := core.GetGoModuleName()
	return map[string]interface{}{
		"Entity":     g.Entity,
		"ModuleName": moduleName,
	}
}

func (g *SchemaGenerator) Generate() (string, error) {
	return RenderTemplateWithFuncMap(
		filepath.Join(g.TemplatePath, g.TemplateFile()),
		g.InputData(),
		map[string]interface{}{
			"joinRules": func(rules []string) string {
				return strings.Join(rules, ",")
			},
		},
	)
}

// RenderTemplateWithFuncMap renders a Go template file with custom funcMap
func RenderTemplateWithFuncMap(templatePath string, data map[string]interface{}, funcMap map[string]interface{}) (string, error) {
	tmpl, err := template.New(filepath.Base(templatePath)).Funcs(template.FuncMap(funcMap)).ParseFiles(templatePath)
	if err != nil {
		return "", err
	}
	var output string
	buf := &outputWriter{str: &output}
	err = tmpl.Execute(buf, data)
	if err != nil {
		return "", err
	}
	return output, nil
}

type outputWriter struct {
	str *string
}

func (w *outputWriter) Write(p []byte) (n int, err error) {
	*w.str += string(p)
	return len(p), nil
}

func JoinPath(elem ...string) string {
	return strings.Join(elem, "/")
}
