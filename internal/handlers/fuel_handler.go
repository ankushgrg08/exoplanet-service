package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/exoplanet-service/internal/models"
	"github.com/exoplanet-service/internal/repository"
	"github.com/exoplanet-service/internal/services"
	"github.com/gorilla/mux"
)

func FuelEstimation(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	exoplanet, found := repository.GetExoplanetByID(id)
	if !found {
		http.Error(w, "Exoplanet not found", http.StatusNotFound)
		return
	}

	crewCapacity, err := strconv.Atoi(r.URL.Query().Get("crewCapacity"))
	if err != nil {
		http.Error(w, "Invalid crew capacity", http.StatusBadRequest)
		return
	}

	fuelCost, err := services.CalculateFuelCost(exoplanet, crewCapacity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := models.FuelEstimationResponse{FuelCost: fuelCost}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
