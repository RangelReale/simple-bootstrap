package simple_bootstrap

import (
	"embed"
	"fmt"
	"html/template"
	"io"
)

var (
	//go:embed templates/*.tmpl
	templatesFS     embed.FS
	parsedTemplates = template.Must(template.ParseFS(templatesFS, "templates/*.tmpl"))
)

func Template(w io.Writer, name string, data any) error {
	return parsedTemplates.ExecuteTemplate(w, fmt.Sprintf("%s.html.tmpl", name), data)
}
