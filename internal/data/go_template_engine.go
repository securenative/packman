package data

import (
	"bytes"
	"text/template"
)

type GoTemplateEngine struct {
}

func NewGoTemplateEngine() *GoTemplateEngine {
	return &GoTemplateEngine{}
}

func (this *GoTemplateEngine) Render(templateText string, data interface{}) (string, error) {
	var out bytes.Buffer
	t := template.New("template")
	t = t.Delims("[[[", "]]]")

	tree, err := t.Parse(templateText)
	if err != nil {
		return "", err
	}

	err = tree.Execute(&out, data)
	if err != nil {
		return "", err
	}

	return out.String(), nil
}
