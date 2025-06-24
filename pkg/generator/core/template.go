package core

import (
	"bytes"
	"os"
	"text/template"
)

// RenderTemplate renders a template file with data and returns the result as string
func RenderTemplate(templatePath string, data interface{}) (string, error) {
	// Read template file
	content, err := os.ReadFile(templatePath)
	if err != nil {
		return "", err
	}

	// Parse template
	tmpl, err := template.New("template").Parse(string(content))
	if err != nil {
		return "", err
	}

	// Execute template
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}
