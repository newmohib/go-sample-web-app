package render

import (
	//"errors"
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/newmohib/go-sample-web-app/html-tamplating/pkg/config"
	"github.com/newmohib/go-sample-web-app/html-tamplating/pkg/models"
)

var functions = template.FuncMap{}

var app *config.AppConfig

// NewTemplate sets the config for the template package
func NewTemplate(a *config.AppConfig) {
	app = a

}

// carate default data
func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// RenderTemplate render template using html/template
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCache {
		// get tehe template cache from the app config
		tc = app.TamplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// tc, err := CreateTemplateCache()

	// if err != nil {
	// 	fmt.Println("error parsing template:", err)
	// 	log.Fatal(err)
	// }

	t, ok := tc[tmpl]

	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	buf := new(bytes.Buffer)
	// check and return td

	td = AddDefaultData(td)

	_ = t.Execute(buf, td)

	_, err := buf.WriteTo(w)

	if err != nil {
		fmt.Println("error parsing template:", err)
		//log.Fatal(err)
	}

	// parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)

	// err1 := parsedTemplate.Execute(w, nil)

	// if err1 != nil {
	// 	fmt.Println("error parsing template:", err1)
	// 	return
	// }
}

// crateTemplateCache creates a template cache as a map

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		fmt.Println("Page is currntly processing", name)
		// we can modify the template into functions here
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)

		if err != nil {
			return myCache, err
		}
		// in tutorial are ./templates/*.layout.tmpl
		matches, err := filepath.Glob("./templates/*.layout.tmpl")

		if err != nil {
			return myCache, err
		}

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
