package repository

import (
	"errors"
	"strconv"
	"sync"

	"github.com/exoplanet-service/internal/models"
)

var (
	exoplanets = make(map[string]models.Exoplanet)
	mutex      = &sync.Mutex{}
	idCounter  = 1
)

func generateID() string {
	mutex.Lock()
	defer mutex.Unlock()
	id := strconv.Itoa(idCounter)
	idCounter++
	return id
}

func AddExoplanet(exoplanet *models.Exoplanet) {
	exoplanet.ID = generateID()
	mutex.Lock()
	exoplanets[exoplanet.ID] = *exoplanet
	mutex.Unlock()
}

func ListExoplanets() []models.Exoplanet {
	mutex.Lock()
	defer mutex.Unlock()
	var result []models.Exoplanet
	for _, exoplanet := range exoplanets {
		result = append(result, exoplanet)
	}
	return result
}

func GetExoplanetByID(id string) (models.Exoplanet, bool) {
	mutex.Lock()
	defer mutex.Unlock()
	exoplanet, found := exoplanets[id]
	return exoplanet, found
}

func UpdateExoplanet(id string, updatedExoplanet models.Exoplanet) error {
	mutex.Lock()
	defer mutex.Unlock()
	if _, found := exoplanets[id]; !found {
		return errors.New("exoplanet not found")
	}
	updatedExoplanet.ID = id
	exoplanets[id] = updatedExoplanet
	return nil
}

func DeleteExoplanet(id string) error {
	mutex.Lock()
	defer mutex.Unlock()
	if _, found := exoplanets[id]; !found {
		return errors.New("exoplanet not found")
	}
	delete(exoplanets, id)
	return nil
}
