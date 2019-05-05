package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGoTemplateEngine_Render(t *testing.T) {

	template := `Helo {{ .Name }} this is test num: {{ .Number }}`
	expected := `Helo World this is test num: 1`
	type data struct {
		Name   string
		Number int
	}

	engine := NewGoTemplateEngine()
	rendered, err := engine.Render(template, data{Name: "World", Number: 1})

	assert.Nil(t, err)
	assert.Equal(t, expected, rendered)
}

func TestGoTemplateEngine_Render_Missing_Arg(t *testing.T) {

	template := `Helo {{ .Name }} this is test num: {{ .Number }}`
	type data struct {
		Name string
	}

	engine := NewGoTemplateEngine()
	rendered, err := engine.Render(template, data{Name: "World"})

	assert.NotNil(t, err)
	assert.Equal(t, "", rendered)
}
