package render

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/WilliamAinsworth/bookings/internal/config"
	"github.com/WilliamAinsworth/bookings/internal/models"
	"github.com/justinas/nosurf"
	"html/template"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

var app *config.AppConfig
var pathToTemplates = "./templates"

// NewTemplate sets the config for the template package
func NewTemplate(a *config.AppConfig) {
	app = a
}

// Add DefaultData adds data for all templates
func AddDefaultData(templateData *models.TemplateData, request *http.Request) *models.TemplateData {
	templateData.Flash = app.Session.PopString(request.Context(), "flash")
	templateData.Error = app.Session.PopString(request.Context(), "error")
	templateData.Warning = app.Session.PopString(request.Context(), "warning")
	templateData.CSRFToken = nosurf.Token(request)
	return templateData
}

// RenderTemplate renders template using http/template
func RenderTemplate(w http.ResponseWriter, request *http.Request, tmpl string, templateData *models.TemplateData) error {
	var templateCache map[string]*template.Template

	// if UseCache is true then read the info from the template cache
	if app.UseCache {
		// get the template cache from the app config
		templateCache = app.TemplateCache
	} else { // otherwise rebuild the template cache
		templateCache, _ = CreateTemplateCache()
	}

	t, ok := templateCache[tmpl]
	if !ok {
		return errors.New("could not get template from cache")
	}

	// holds bytes (parsed template)
	buf := new(bytes.Buffer)

	// add data that should be available to all pages
	templateData = AddDefaultData(templateData, request)

	// renders the page
	_ = t.Execute(buf, templateData)

	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser")
		return err
	}

	return nil

	//parsedTemplate, _ :=  template.ParseFiles("./templates/" + tmpl)
	//err = parsedTemplate.Execute(w, nil)
	//if err != nil {
	//	fmt.Println("error parsing template:", err)
	//	return
	//}
}

// CreateTemplateCache creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {
	// create a map with key: string -- value: pointer to template
	myCache := map[string]*template.Template{}

	// find all pages in the templates folder
	pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.tmpl", pathToTemplates))
	if err != nil {
		return myCache, err
	}

	// loop through all found pages and print out name of current page
	for _, page := range pages {
		name := filepath.Base(page)

		templateSet, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			templateSet, err = templateSet.ParseGlob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = templateSet
	}
	return myCache, nil
}
