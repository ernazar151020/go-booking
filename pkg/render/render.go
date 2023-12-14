package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/ernazar151020/go-packages/config"
	"github.com/ernazar151020/go-packages/models"
)

var functions = template.FuncMap{}

var app *config.AppConfig

// setthe config  for thetemplate page
func NewTemplates(a *config.AppConfig) {
	app = a
}

func RenderTemplate(res http.ResponseWriter, file string, td *models.TemplateData) {

	var tc map[string]*template.Template

	if app.UseCache {

		// Get the template cahce from config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// tc , err := CreateTemplateCache()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	t, ok := tc[file]

	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	_ = t.Execute(buf, td)

	_, err := buf.WriteTo(res)

	if err != nil {
		fmt.Println("Error writing template to browser", err)
	}
	// parsedTemplate, _ := template.ParseFiles("./templates/" + file)
	// err = parsedTemplate.Execute(res, nil)
	// if err != nil {
	// 	fmt.Println("Error executing template : ", err)
	// }
	// return
}

func CreateTemplateCache() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.pages.html")
	if err != nil {

		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		fmt.Println("Page is currenlty", page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*layout.html")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*layout.html")
			if err != nil {
				return myCache, err
			}
			myCache[name] = ts

		}

	}

	return myCache, nil

}
