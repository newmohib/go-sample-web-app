package main

import (
	//"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/newmohib/go-sample-web-app/html-tamplating/pkg/config"
	"github.com/newmohib/go-sample-web-app/html-tamplating/pkg/handlers"
	"github.com/newmohib/go-sample-web-app/html-tamplating/render"
)

// application portNumber
const portNumber = ":8080"

// main is the main application function
func main() {
	var app config.AppConfig
	// get the template cache from the carate app config
	tc, err := render.CreateTemplateCache()

	if err != nil {
		fmt.Println("error parsing template:", err)
		log.Fatal("Can not create template cache")
	}
	app.TamplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo((&app))
	handlers.NewHandlers(repo)

	render.NewTemplate(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println("Starting application on port:", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
