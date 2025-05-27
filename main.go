// main.go
package main

import (
	"log"
	"net/http"
	"url-shortener/router"
)

func main() {
	r := router.NewRouter()
	log.Println("Starting server at :8080")
	http.ListenAndServe(":8080", r)
}
