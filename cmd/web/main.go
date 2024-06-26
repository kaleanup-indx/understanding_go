package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kaleanup-indx/m/v2/pkg/config"
	"github.com/kaleanup-indx/m/v2/pkg/handlers"
	"github.com/kaleanup-indx/m/v2/pkg/render"
)

const portNumber = ":8080"

// main the main function of the applicaiton.
func main() {

	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Printf((fmt.Sprintf("Starting application on port %s \n", portNumber)))

	srv := http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)

}
