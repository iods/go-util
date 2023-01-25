package strutil

import (
	"bytes"
	"text/template"
)

// Template indicates a struct for different output types.
type Template struct {
	*template.Template
}

func Exec(f string, d interface{}) ([]byte, error) {
	t, err := Parse(f)
	if err != nil {
		return nil, err
	}
	return t.Exec(d)
}

// Parse returns a formatted template.
func Parse(s string) (*Template, error) {
	t, err := template.New("_").Parse(s)
	if err != nil {
		return nil, err
	}
	return &Template{t}, nil
}

// Exec is the intern
func (t *Template) Exec(d interface{}) ([]byte, error) {
	b := &bytes.Buffer{}
	err := t.Execute(b, d)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
