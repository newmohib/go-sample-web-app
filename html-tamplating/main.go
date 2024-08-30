package main

import (
	//"errors"
	"fmt"
	"net/http"
)

// application portNumber
const portNumber = ":8080"

// main is the main application function
func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println("Starting application on port:", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
