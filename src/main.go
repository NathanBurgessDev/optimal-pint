package main

import (
	"net/http"
	"optimal-pint/src/internal/service"
)

func main() {
	http.HandleFunc("/hello", service.Hello)

	http.ListenAndServe(":8080", nil)
}
