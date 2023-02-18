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

// variable de tipo puntero a la configuracion de la aplicacion
var app *config.AppConfig

// NewTemplates establece la configuracion para el paquete de plantilla
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(data *models.TemplatesData) *models.TemplatesData {
	return data
}

// RenderTemplate renderiza una plantilla
func RenderTemplate(w http.ResponseWriter, tempName string, data *models.TemplatesData) {
	var tempsCache map[string]*template.Template

	// toma el valor de la configuracion de la aplicacion y si se establece que se esta utilizando el cache entonces sencillamente lo usa y si no entonces lee desde el disco cada vez que ingrese a una apgina desde el servidor, esto es util para cuando se esta en desarrollo
	if app.UseCache {
		tempsCache = app.TempsCache
	} else {
		tempsCache, _ = CreateTemplatesCache()
	}

	// extrae de el cache la plantilla a la que se esta accediendo desde el servidor
	temp, ok := tempsCache[tempName]
	if !ok {
		log.Fatal("error al extraer plantilla")
	}

	// crea un nuevo buffer
	buff := new(bytes.Buffer)

	data = AddDefaultData(data)

	err := temp.Execute(buff, data)
	if err != nil {
		log.Println(err)
	}

	_, err = buff.WriteTo(w)
	if err != nil {
		log.Println("error al escribir la plantilla en el navegador", err)
	}
}

// CreateTemplatesCache crea un cache de plantillas, como mapa
func CreateTemplatesCache() (map[string]*template.Template, error) {
	tempsCache := make(map[string]*template.Template)

	// retorna un slice con la ruta completa de los archivos terminados en *_page.gohtml
	temps, err := filepath.Glob("./templates/*_page.gohtml")
	if err != nil {
		return tempsCache, err
	}

	// recorre todo el slice anterior 
	for _, page := range temps {

		// devuelve solo la base de la ruta sin ningun separador
		tempName := filepath.Base(page)
		// [New] asigna una nueva plantilla con el nombre que se le pase. [ParseFiles] se encarga de analizar la plantilla
		temp, err := template.New(tempName).ParseFiles(page)
		if err != nil {
			return tempsCache, err
		}

		// retorna un slice con la ruta completa de los archivos terminados en *_layout.gohtml
		layouts, err := filepath.Glob("./templates/*_layout.gohtml")
		if err != nil {
			return tempsCache, err
		}

		// esto se realiza para determinar si hay al menos un archivo dentro del slice anterior, osea un layout
		if len(layouts) > 0 {

			// Si encontro layouts, entonces asocia a cada plantilla esos layouts y los analiza todoo devolviendo por ultimo una sola platilla.
			temp, err = temp.ParseGlob("./templates/*_layout.gohtml")
			if err != nil {
				return tempsCache, err
			}
		}
		// agrega al cache de plantillas la nueva platilla ya analizada unida con los layouts igualmente analizados, y como llave pone el nombre que seria la base de la plantilla.
		tempsCache[tempName] = temp
	}
	// si todo salio bien devuelve el cache de plantillas y un error nulo osea que no existe
	return tempsCache, nil
}
