package domain

// Location represents a geographical location
type Location struct {
	Latitude  float64
	Longitude float64
}

// NewLocation creates a new Location instance
func NewLocation(latitude, longitude float64) *Location {
	return &Location{
		Latitude:  latitude,
		Longitude: longitude,
	}
}

