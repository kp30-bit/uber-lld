package strategy

import (
	"fmt"

	"github.com/kamalpratik/Uber-Ola-Low-Level-Design/internal/domain"
	"github.com/kamalpratik/Uber-Ola-Low-Level-Design/pkg/interfaces"
)

// DefaultPricingStrategy implements the default pricing strategy
type DefaultPricingStrategy struct{}

// NewDefaultPricingStrategy creates a new DefaultPricingStrategy instance
func NewDefaultPricingStrategy() interfaces.PricingStrategy {
	return &DefaultPricingStrategy{}
}

// CalculatePrice calculates the price using the default strategy
func (d *DefaultPricingStrategy) CalculatePrice(tripMetaData *domain.TripMetaData) float64 {
	fmt.Println("Based on default strategy, price is 100")
	return 100.0
}
