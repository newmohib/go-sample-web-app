package main

import (
	"errors"
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

// Create a Divided handler
func Divided(w http.ResponseWriter, r *http.Request) {

	f, err := dividedValues(100.0, 0.0)

	if err != nil {
		fmt.Fprintf(w, "%s :", err)
		return
	}

	_, _ = fmt.Fprintf(w, "%f divided by %f is: %f", 100.0, 0.0, f)

}

// Create a Divided handler
func dividedValues(x, y float32) (float32, error) {

	if y <= 0 {
		err := errors.New("can not divide by zero")
		return 0, err
	}

	return x / y, nil

}

// main is the main application function
func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divided", Divided)

	fmt.Println("Starting application on port:", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
