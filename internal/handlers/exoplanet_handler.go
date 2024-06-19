package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/exoplanet-service/internal/models"
	"github.com/exoplanet-service/internal/repository"
	"github.com/exoplanet-service/internal/validation"
	"github.com/gorilla/mux"
)

func AddExoplanet(w http.ResponseWriter, r *http.Request) {
	var exoplanet models.Exoplanet
	err := json.NewDecoder(r.Body).Decode(&exoplanet)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := validation.ValidateExoplanet(exoplanet); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	repository.AddExoplanet(&exoplanet)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(exoplanet)
}

func ListExoplanets(w http.ResponseWriter, r *http.Request) {
	exoplanets := repository.ListExoplanets()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(exoplanets)
}

func GetExoplanetByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	exoplanet, found := repository.GetExoplanetByID(id)
	if !found {
		http.Error(w, "Exoplanet not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(exoplanet)
}

func UpdateExoplanet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var updatedExoplanet models.Exoplanet
	err := json.NewDecoder(r.Body).Decode(&updatedExoplanet)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := validation.ValidateExoplanet(updatedExoplanet); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := repository.UpdateExoplanet(id, updatedExoplanet); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedExoplanet)
}

func DeleteExoplanet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	if err := repository.DeleteExoplanet(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
