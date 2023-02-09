package main

import (
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/stivn1/bookings/pkg/config"
	"github.com/stivn1/bookings/pkg/handlers"
	"github.com/stivn1/bookings/pkg/render"
)

const PORTNUMBER = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tempsCache, err := render.CreateTemplatesCache()
	if err != nil {
		log.Fatal("no se pudo crear el cache")
	}

	app.TempsCache = tempsCache
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	server := &http.Server{
		Addr: PORTNUMBER,
		Handler: routes(&app),
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal("no se pudo cargar el servidor")
	}
}
