package services

import (
	"errors"

	"github.com/exoplanet-service/internal/models"
)

func CalculateGravity(exoplanet models.Exoplanet) (float64, error) {
	switch exoplanet.Type {
	case "GasGiant":
		return 0.5 / (exoplanet.Radius * exoplanet.Radius), nil
	case "Terrestrial":
		return exoplanet.Mass / (exoplanet.Radius * exoplanet.Radius), nil
	default:
		return 0, errors.New("invalid exoplanet type")
	}
}

func CalculateFuelCost(exoplanet models.Exoplanet, crewCapacity int) (float64, error) {
	gravity, err := CalculateGravity(exoplanet)
	if err != nil {
		return 0, err
	}
	return float64(exoplanet.Distance) / (gravity * gravity) * float64(crewCapacity), nil
}
