package main

import (
	"fmt"
	"net/http"
)

// application portNumber
const portNumber = ":8080"

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("This is the home page")

	fmt.Fprintf(w, "This is the home page")

	// n, err := fmt.Fprintf(w, "Hello, World!")

	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(fmt.Sprintf("Number of bytes written: %d", n))

}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(3, 2)
	fmt.Println("This is the about page")

	_, _ = fmt.Fprintf(w, "This is the about page and the sum is: %d", sum)

}

// addAalues  adds two integers and retun thw sum
func addValues(x int, y int) int {
	if x <= 0 || y <= 0 {
		return 0
	}

	sum := x + y
	return sum

}

// main is the main application function

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println("Starting application on port:", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
