package main

import (
	//"errors"
	"fmt"
	"net/http"

	"github.com/newmohib/go-sample-web-app/html-tamplating/pkg/handlers"
)

// application portNumber
const portNumber = ":8080"

// main is the main application function
func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println("Starting application on port:", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
