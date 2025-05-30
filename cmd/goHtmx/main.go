package main

import (
	"html/template"
	"io"
	"os"

	"github.com/Younes-khadraoui/Go-HTMX/internal/router"
	"github.com/labstack/echo/v4"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	err := t.templates.ExecuteTemplate(w, name, data)
	if err != nil {
		c.Logger().Errorf("Render error: %v", err)
	}
	return err
}

func main() {
	e := echo.New()

	// Load templates
	tmpl := template.Must(template.ParseGlob("web/templates/**/*.html"))
	e.Renderer = &TemplateRenderer{templates: tmpl}

	// Static files (e.g. Tailwind, Alpine)
	e.Static("/static", "web/static")

	// Register routes
	router.RegisterRoutes(e)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	e.Logger.Fatal(e.Start(":" + port))
}
