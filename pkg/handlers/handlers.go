package handlers

import (
	"net/http"

	"github.com/stivn1/bookings/pkg/config"
	"github.com/stivn1/bookings/pkg/models"
	"github.com/stivn1/bookings/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home_page.gohtml", &models.TemplatesData{})
}
