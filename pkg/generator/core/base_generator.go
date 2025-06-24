package core

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// Generator interface chuẩn hóa cho các generator layer
// (repository, service, handler, ...)
type Generator interface {
	TemplateFile() string
	OutputFilename() string
	ComponentFolder() string
	InputData() map[string]interface{}
	Generate() (string, error) // render nội dung file
}

// BaseGenerator struct dùng để embed cho các generator cụ thể
// Chứa logic render template và các trường chung
// Các generator cụ thể chỉ cần implement các method abstract
// và có thể override InputData nếu muốn

type BaseGenerator struct {
	Entity       *Entity
	TemplatePath string
	Force        bool
	// ... các trường khác nếu cần
}

// RenderTemplate renders a Go template file with data and returns the result as string
func (g *BaseGenerator) RenderTemplate(templateFile string, data map[string]interface{}) (string, error) {
	tmplPath := filepath.Join(g.TemplatePath, templateFile)
	tmpl, err := template.New(filepath.Base(templateFile)).Funcs(template.FuncMap{
		// custom funcs nếu cần
	}).ParseFiles(tmplPath)
	if err != nil {
		return "", err
	}
	var output string
	buf := &outputWriter{&output}
	err = tmpl.Execute(buf, data)
	if err != nil {
		return "", err
	}
	return output, nil
}

// outputWriter implements io.Writer for string
// (để ghi template ra string thay vì file)
type outputWriter struct {
	str *string
}

func (w *outputWriter) Write(p []byte) (n int, err error) {
	*w.str += string(p)
	return len(p), nil
}

// Hàm tạo thư mục nếu chưa tồn tại
func EnsureDir(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return os.MkdirAll(dir, 0755)
	}
	return nil
}

// GetGoModuleName reads go.mod and returns the module name
func GetGoModuleName() (string, error) {
	f, err := os.Open("go.mod")
	if err != nil {
		return "", err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "module ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "module ")), nil
		}
	}
	return "", nil
}
