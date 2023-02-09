package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/stivn1/bookings/pkg/config"
	"github.com/stivn1/bookings/pkg/models"
)

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(data *models.TemplatesData) *models.TemplatesData {
	return data
}

func RenderTemplate(w http.ResponseWriter, tempName string, data *models.TemplatesData) {
	var tempsCache map[string]*template.Template

	if app.UseCache {
		tempsCache = app.TempsCache
	} else {
		tempsCache, _ = CreateTemplatesCache()
	}

	temp, ok := tempsCache[tempName]
	if !ok {
		log.Fatal("error al extraer plantilla")
	}

	buff := new(bytes.Buffer)

	data = AddDefaultData(data)

	err := temp.Execute(buff, data)
	if err != nil {
		log.Println(err)
	}

	_, err = buff.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplatesCache() (map[string]*template.Template, error) {
	tempsCache := make(map[string]*template.Template)

	temps, err := filepath.Glob("./templates/*_page.gohtml")
	if err != nil {
		return tempsCache, err
	}

	for _, page := range temps {
		tempName := filepath.Base(page)
		temp, err := template.New(tempName).ParseFiles(page)
		if err != nil {
			return tempsCache, err
		}

		layouts, err := filepath.Glob("./templates/*_layout.gohtml")
		if err != nil {
			return tempsCache, err
		}

		if len(layouts) > 0 {
			temp, err = temp.ParseGlob("./templates/*_layout.gohtml")
			if err != nil {
				return tempsCache, err
			}
		}
		tempsCache[tempName] = temp
	}
	return tempsCache, nil
}
