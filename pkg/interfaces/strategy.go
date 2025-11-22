package interfaces

import "github.com/kamalpratik/Uber-Ola-Low-Level-Design/internal/domain"

// PricingStrategy defines the interface for pricing strategies
type PricingStrategy interface {
	CalculatePrice(tripMetaData *domain.TripMetaData) float64
}

// DriverMatchingStrategy defines the interface for driver matching strategies
type DriverMatchingStrategy interface {
	MatchDriver(tripMetaData *domain.TripMetaData) *domain.Driver
}

