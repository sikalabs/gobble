package template_utils

import (
	"strings"
	"text/template"
)

func RenderTemplateToString(
	templateString, templateName string,
	data interface{},
) (string, error) {
	t, err := template.New(templateName).Parse(templateString)
	if err != nil {
		return "", err
	}
	var buf strings.Builder
	err = t.Execute(&buf, data)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
