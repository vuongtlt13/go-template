package backend

import (
	"path/filepath"
	"yourapp/pkg/generator/core"
)

type HandlerGenerator struct {
	*core.BaseGenerator
}

func NewHandlerGenerator(entity *core.Entity, templatePath string, force bool) core.Generator {
	return &HandlerGenerator{
		BaseGenerator: &core.BaseGenerator{
			Entity:       entity,
			TemplatePath: templatePath,
			Force:        force,
		},
	}
}

func (g *HandlerGenerator) TemplateFile() string { return "backend/handler.go.tmpl" }
func (g *HandlerGenerator) OutputFilename() string {
	return g.Entity.ModelTextForm.ModelNameSnakeCase + ".go"
}
func (g *HandlerGenerator) ComponentFolder() string { return "internal/handler/admin" }
func (g *HandlerGenerator) InputData() map[string]interface{} {
	moduleName, _ := core.GetGoModuleName()
	return map[string]interface{}{
		"Entity":     g.Entity,
		"ModuleName": moduleName,
		"Text":       g.Entity.ModelTextForm,
	}
}

func (g *HandlerGenerator) Generate() (string, error) {
	return RenderTemplateWithFuncMap(
		filepath.Join(g.TemplatePath, g.TemplateFile()),
		g.InputData(),
		nil, // Không cần custom funcMap cho handler
	)
}
