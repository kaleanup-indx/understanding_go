package main

import (
	"fmt"
	"net/http"

	"github.com/kaleanup-indx/m/v2/pkg/handlers"
)

const portNumber = ":8080"

// main the main function of the applicaiton.
func main() {

	http.HandleFunc(
		"/",
		handlers.Home,
	)

	http.HandleFunc(
		"/about",
		handlers.About,
	)

	fmt.Printf((fmt.Sprintf("Starting application on port %s", portNumber)))
	http.ListenAndServe(portNumber, nil)
}
