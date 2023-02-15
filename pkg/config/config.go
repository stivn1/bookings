package config

import (
	"html/template"

	"github.com/alexedwards/scs/v2"
)

// AppConfig estructura que contiene todas las configuraciones de la aplicacion
type AppConfig struct {
	UseCache     bool
	TempsCache   map[string]*template.Template
	InProduction bool
	Session      *scs.SessionManager
}
