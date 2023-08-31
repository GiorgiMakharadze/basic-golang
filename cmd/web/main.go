package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/GiorgiMakharadze/basic-golang/pkg/config"
	"github.com/GiorgiMakharadze/basic-golang/pkg/handlers"
	"github.com/GiorgiMakharadze/basic-golang/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"
var app config.AppConfig
var session *scs.SessionManager


func main() {

	// In production set InProduction to true
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot crate template cache",err)
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	srv := &http.Server {
		Addr: portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
