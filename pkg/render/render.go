package render

import (
	"bytes"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

// Renders template using html/template
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	// create a template cache
	tc, err := createTemplateCache()
	if err != nil {
		log.Fatal(err)
	}
	// get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal(err)
	}
	// render template

	// create buffer
	buf := new(bytes.Buffer)

	// write template to buffer
	err = t.Execute(buf, nil)
	if err != nil {
		log.Println(err)
	}
	// write buffer to w (http.ResponseWriter)
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func createTemplateCache() (map[string]*template.Template, error) {
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
