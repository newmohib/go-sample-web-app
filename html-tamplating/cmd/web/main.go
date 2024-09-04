package main

import (
	//"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/newmohib/go-sample-web-app/html-tamplating/pkg/config"
	"github.com/newmohib/go-sample-web-app/html-tamplating/pkg/handlers"
	"github.com/newmohib/go-sample-web-app/html-tamplating/render"
)

// application portNumber
const portNumber = ":8080"

// initialize app config
// its alos using into middleware or any others into main package
var app config.AppConfig
// initialize sessin manager and its alos using into middleware or any others into main package
var session *scs.SessionManager

// main is the main application function
func main() {

	// change this to true when in production
	app.InProduction = false

	// set session
	session = scs.New()
	session.Lifetime = 24 * time.Hour // session will expire after 24 hours
	session.Cookie.Persist = true     // session will expire after the browser is closed
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	// store thsis session into app config
	app.Session = session

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

	fmt.Println("Starting application on port:", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}
