package backend

import (
	"path/filepath"
	"yourapp/pkg/generator/core"
)

type ServiceGenerator struct {
	*core.BaseGenerator
}

func NewServiceGenerator(entity *core.Entity, templatePath string, force bool) core.Generator {
	return &ServiceGenerator{
		BaseGenerator: &core.BaseGenerator{
			Entity:       entity,
			TemplatePath: templatePath,
			Force:        force,
		},
	}
}

func (g *ServiceGenerator) TemplateFile() string { return "backend/service.go.tmpl" }
func (g *ServiceGenerator) OutputFilename() string {
	return g.Entity.ModelTextForm.ModelNameSnakeCase + ".go"
}
func (g *ServiceGenerator) ComponentFolder() string { return "internal/service" }
func (g *ServiceGenerator) InputData() map[string]interface{} {
	moduleName, _ := core.GetGoModuleName()
	return map[string]interface{}{
		"Entity":     g.Entity,
		"ModuleName": moduleName,
		"Text":       g.Entity.ModelTextForm,
	}
}

func (g *ServiceGenerator) Generate() (string, error) {
	return RenderTemplateWithFuncMap(
		filepath.Join(g.TemplatePath, g.TemplateFile()),
		g.InputData(),
		nil, // Không cần custom funcMap cho service
	)
}
