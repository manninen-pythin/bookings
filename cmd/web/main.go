package main

import (
	"fmt"
	"net/http"

	"github.com/manninen-pythin/bookings/pkg/handlers"
)

const portNumber = ":8080"

// main app function
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("starting application on port%s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
