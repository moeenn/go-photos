package templates

import (
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
	"log"
)

type Renderer struct {
	templates *template.Template
}

func (r *Renderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	err := r.templates.ExecuteTemplate(w, name, data)
	if err != nil {
		log.Printf("[error] %s", err)
	}

	return err
}

func New(path string) *Renderer {
	return &Renderer{
		templates: template.Must(template.ParseGlob(path)),
	}
}
