package models

type Exoplanet struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Distance    int     `json:"distance"`   // light years
	Radius      float64 `json:"radius"`     // Earth-radius units
	Mass        float64 `json:"mass,omitempty"` // Earth-mass units, only for Terrestrial
	Type        string  `json:"type"`       // "GasGiant" or "Terrestrial"
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type FuelEstimationResponse struct {
	FuelCost float64 `json:"fuelCost"`
}
