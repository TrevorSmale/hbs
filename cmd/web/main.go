package main

import (
	"fmt"
	"net/http"

	"github.com/trevorsmale/hbs/pkg/handlers"
)

const portNumber = ":8080"

// mian is the main application function
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprint("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
