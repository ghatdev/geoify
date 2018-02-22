package main

import (
	"geoify/api"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func App() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/{ip}", api.GetIPGeoInfo)

	r.MethodNotAllowedHandler = http.HandlerFunc(api.MethodNotAllowed)
	r.NotFoundHandler = http.HandlerFunc(api.NotFound)

	return cors.Default().Handler(r)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Fatal(http.ListenAndServe(":"+port, App()))
}
