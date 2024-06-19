package main

import (
	"log"
	"net/http"

	"github.com/exoplanet-service/internal/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/exoplanets", handlers.AddExoplanet).Methods("POST")
	r.HandleFunc("/exoplanets", handlers.ListExoplanets).Methods("GET")
	r.HandleFunc("/exoplanets/{id}", handlers.GetExoplanetByID).Methods("GET")
	r.HandleFunc("/exoplanets/{id}", handlers.UpdateExoplanet).Methods("PUT")
	r.HandleFunc("/exoplanets/{id}", handlers.DeleteExoplanet).Methods("DELETE")
	r.HandleFunc("/exoplanets/{id}/fuel", handlers.FuelEstimation).Methods("GET")

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
