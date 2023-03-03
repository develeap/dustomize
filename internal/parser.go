// Package internal provides internal helper functions.
package internal

import (
	"bytes"
	"text/template"

	"github.com/Masterminds/sprig/v3"
)

// Parses a file (origin) with a config (data)
func ParseFile(templateName string, text string, configData map[interface{}]interface{}) (string, error) {
	t, err := template.New(templateName).Funcs(sprig.FuncMap()).Parse(text)
	if err != nil {
		return "", err
	}

	var buffer bytes.Buffer
	err = t.Execute(&buffer, configData)

	if err != nil {
		return "", err
	}

	return buffer.String(), nil
}
