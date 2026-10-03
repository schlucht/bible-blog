package frontend

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/schlucht/bibleblog/cmd/web/config"
)

type TemplateData struct {
	StringMap       map[string]string
	IntMap          map[string]int
	FloatMap        map[string]float32
	Data            map[string]interface{}
	CSRFToken       string
	Flash           string
	Warning         string
	Error           string
	IsAuthenticated int
	API             string
}

const (
	tmpl = "templates/"
)

var app *config.AppConfig

func NewRenderer(a *config.AppConfig) {
	app = a
}

var defaultPartials = []string{"header"}
var functions = template.FuncMap{}

//go:embed templates
var templateFS embed.FS

func AddDefaultData(td *TemplateData, r *http.Request) *TemplateData {

	return td
}

func parseLayouts() ([]string, error) {
	var paths []string
	err := fs.WalkDir(templateFS, "templates/layouts", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && strings.HasSuffix(path, ".tmpl") {
			paths = append(paths, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	if len(paths) == 0 {
		paths = append(paths, "templates/layouts/base.layout.tmpl")
	}
	return paths, nil
}

func parseTemplate(partials []string, page, templateToRender string) (*template.Template, error) {
	var t *template.Template
	var err error

	if len(partials) > 0 {
		for i, x := range partials {
			partials[i] = fmt.Sprintf("templates/partials/%s.partial.tmpl", x)
		}
	}

	layouts, err := parseLayouts()
	if err != nil {
		app.ErrorLog.Println("Find no layouts", err)
		return nil, err
	}
	allFiles := []string{}
	allFiles = append(allFiles, partials...)
	allFiles = append(allFiles, templateToRender)
	allFiles = append(allFiles, layouts...)

	t, err = template.New(filepath.Base(templateToRender)).
		Funcs(functions).
		ParseFS(templateFS, allFiles...)
	if err != nil {
		app.ErrorLog.Println("Error parsing templates files", err)
		return nil, err
	}
	app.TemplateCache[templateToRender] = t
	return t, nil
}

func RenderTemplate(w http.ResponseWriter, r *http.Request, page string, td *TemplateData, partials ...string) error {
	var t *template.Template
	var err error
	partials = append(partials, defaultPartials...)
	templateToRender := fmt.Sprintf("templates/pages/%s.page.tmpl", page)

	_, templateInMap := app.TemplateCache[templateToRender]

	if templateInMap {
		t = app.TemplateCache[templateToRender]
	} else {
		t, err = parseTemplate(partials, page, templateToRender)
		if err != nil {
			app.ErrorLog.Println(err)
			return err
		}
	}
	if td == nil {
		td = &TemplateData{}
	}
	td = AddDefaultData(td, r)
	err = t.Execute(w, td)
	if err != nil {
		app.ErrorLog.Println(err)
		return err
	}
	return nil
}
