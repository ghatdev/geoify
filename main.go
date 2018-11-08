package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ghatdev/geoify/api"

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

func init() {
	log.Println("Downloading DB file...")
	updateDB()
	log.Println("DB file successfully updated!")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	addCron()

	api.OpenDB()

	log.Fatal(http.ListenAndServe(":"+port, App()))
}
