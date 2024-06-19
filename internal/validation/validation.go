package validation

import (
	"errors"

	"github.com/exoplanet-service/internal/models"
)

func ValidateExoplanet(exoplanet models.Exoplanet) error {
	if exoplanet.Distance <= 10 || exoplanet.Distance >= 1000 {
		return errors.New("distance must be between 10 and 1000 light years")
	}
	if exoplanet.Radius <= 0.1 || exoplanet.Radius >= 10 {
		return errors.New("radius must be between 0.1 and 10 Earth-radius units")
	}
	if exoplanet.Type != "GasGiant" && exoplanet.Type != "Terrestrial" {
		return errors.New("type must be either 'GasGiant' or 'Terrestrial'")
	}
	if exoplanet.Type == "Terrestrial" {
		if exoplanet.Mass <= 0.1 || exoplanet.Mass >= 10 {
			return errors.New("mass must be between 0.1 and 10 Earth-mass units for Terrestrial exoplanets")
		}
	}
	return nil
}
