package backend

import (
	"yourapp/pkg/generator/core"
)

type RepositoryGenerator struct {
	*core.BaseGenerator
}

func NewRepositoryGenerator(entity *core.Entity, templatePath string, force bool) core.Generator {
	return &RepositoryGenerator{
		BaseGenerator: &core.BaseGenerator{
			Entity:       entity,
			TemplatePath: templatePath,
			Force:        force,
		},
	}
}

func (g *RepositoryGenerator) TemplateFile() string { return "backend/repository.go.tmpl" }
func (g *RepositoryGenerator) OutputFilename() string {
	return g.Entity.ModelTextForm.ModelNameSnakeCase + ".go"
}
func (g *RepositoryGenerator) ComponentFolder() string { return "internal/repository" }
func (g *RepositoryGenerator) InputData() map[string]interface{} {
	moduleName, _ := core.GetGoModuleName()
	return map[string]interface{}{
		"Entity":     g.Entity,
		"ModuleName": moduleName,
	}
}
func (g *RepositoryGenerator) Generate() (string, error) {
	return g.RenderTemplate(g.TemplateFile(), g.InputData())
}
