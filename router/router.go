package router

import (
	"github.com/gorilla/mux"
	"url-shortener/handler"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/shorten", handler.ShortenURL).Methods("POST")
	r.HandleFunc("/metrics/top-domains", handler.GetTopDomains).Methods("GET")
	r.HandleFunc("/{shortID}", handler.Redirect).Methods("GET")

	return r
}
