package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/manninen-pythin/bookings/pkg/config"
	"github.com/manninen-pythin/bookings/pkg/handlers"
	"github.com/manninen-pythin/bookings/pkg/render"
)

const portNumber = ":8080"

// main app function
func main() {
	var app config.AppConfig
	// create a template cache
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}
	// creates a new Config struct and stores the cached templates in .TemplateCache
	app.TemplateCache = tc
	// allows us to set whether the application uses cache or reads templates from disk if set to false (dev mode)
	app.UseCache = false

	// creates a repository which is a struct
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("starting application on port%s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
