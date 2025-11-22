package strategy

import (
	"fmt"

	"github.com/kamalpratik/Uber-Ola-Low-Level-Design/internal/domain"
	"github.com/kamalpratik/Uber-Ola-Low-Level-Design/pkg/common"
	"github.com/kamalpratik/Uber-Ola-Low-Level-Design/pkg/interfaces"
)

// RatingBasedPricingStrategy implements pricing based on rider rating
type RatingBasedPricingStrategy struct{}

// NewRatingBasedPricingStrategy creates a new RatingBasedPricingStrategy instance
func NewRatingBasedPricingStrategy() interfaces.PricingStrategy {
	return &RatingBasedPricingStrategy{}
}

// CalculatePrice calculates the price based on rider rating
func (r *RatingBasedPricingStrategy) CalculatePrice(tripMetaData *domain.TripMetaData) float64 {
	var price float64
	if common.IsHighRating(tripMetaData.GetRiderRating()) {
		price = 55.0
	} else {
		price = 65.0
	}
	fmt.Printf("Based on %s rider rating, price of the trip is %.2f\n",
		common.RatingToString(tripMetaData.GetRiderRating()), price)
	return price
}
