package main

import (
	"linknau-test/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/login", handlers.Login)
	http.HandleFunc("/data", handlers.Authenticate(handlers.FetchData))

	http.ListenAndServe(":8080", nil)
}
