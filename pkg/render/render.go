package render

import (
	"bytes"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/manninen-pythin/bookings/pkg/config"
	"github.com/manninen-pythin/bookings/pkg/models"
)

// Declare a AppConfig Struct
var app *config.AppConfig

// Store Cache in App
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// Renders template using html/template
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template

	if app.UseCache {
		tc = app.TemplateCache
	} else {
		// Get template cache from the app config
		tc, _ = CreateTemplateCache()
	}

	// get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	// create buffer
	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	// write template to buffer
	_ = t.Execute(buf, td)
	// write buffer to w (http.ResponseWriter)
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	// range through all files ending with page.tmpl from ./templates
	for _, page := range pages {
		// get file name without path
		name := filepath.Base(page)
		// creates new template, and the reads the template stored in page into it and stores the value in ts
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		// gets list of layout pages
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}
		// if there are layouts add them to ts
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
