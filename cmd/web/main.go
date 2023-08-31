package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/GiorgiMakharadze/basic-golang/pkg/config"
	"github.com/GiorgiMakharadze/basic-golang/pkg/handlers"
	"github.com/GiorgiMakharadze/basic-golang/pkg/render"
)

const portNumber = ":8080"


func main() {
	var app config.AppConfig

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
