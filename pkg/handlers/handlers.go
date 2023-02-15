package handlers

import (
	"net/http"

	"github.com/stivn1/bookings/pkg/config"
	"github.com/stivn1/bookings/pkg/models"
	"github.com/stivn1/bookings/pkg/render"
)

// Repo es el repositorio usado por los controladores
var Repo *Repository

// Repository estructura que contiene una dato de tipo AppConfig con toda la configuracion de la aplicacion
type Repository struct {
	App *config.AppConfig
}

// NewRepo funcion que crea un nuevo repositorio
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers establece el repositorio para los controladores
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home_page.gohtml", &models.TemplatesData{})
}
