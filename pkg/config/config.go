package config

import "html/template"

type AppConfig struct {
	UseCache     bool
	TempsCache   map[string]*template.Template
	InProduction bool
}
