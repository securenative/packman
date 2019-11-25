package data

import (
	"bytes"
	"github.com/securenative/packman/internal/etc"
	"text/template"
)

type golangTemplateEngine struct {
}

func NewGolangTemplateEngine() *golangTemplateEngine {
	return &golangTemplateEngine{}
}

func (this *golangTemplateEngine) Run(filePath string, data map[string]interface{}) error {
	var out bytes.Buffer
	t := template.New("template")
	t.Delims("{{{", "}}}")

	etc.PrintInfo("Trying to template the following file: %s...", filePath)
	templateText, err := etc.ReadFile(filePath)

	tree, err := t.Parse(templateText)
	if err != nil {
		return err
	}

	err = tree.Execute(&out, data)
	if err != nil {
		return err
	}

	err = etc.WriteFile(filePath, out.String(), etc.StringEncoder)
	if err != nil {
		return nil
	}

	etc.PrintSuccess("	OK\n")
	return nil
}
