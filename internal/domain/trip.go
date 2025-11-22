package domain

import (
	"fmt"
	"github.com/kamalpratik/Uber-Ola-Low-Level-Design/pkg/common"
)

// Trip represents a trip in the system
type Trip struct {
	rider                 *Rider
	driver                *Driver
	srcLoc                *Location
	dstLoc                *Location
	status                common.TripStatus
	tripID                int
	price                 float64
	pricingStrategy       interface{} // Stored but not used, kept for compatibility with C++ design
	driverMatchingStrategy interface{} // Stored but not used, kept for compatibility with C++ design
}

// NewTrip creates a new Trip instance
// This is not threadsafe for tripId generation and is just for demo purposes
func NewTrip(
	rider *Rider,
	driver *Driver,
	srcLoc *Location,
	dstLoc *Location,
	price float64,
	pricingStrategy interface{},
	driverMatchingStrategy interface{},
) *Trip {
	tripID := common.NextTripID
	common.NextTripID++

	return &Trip{
		rider:                 rider,
		driver:                driver,
		srcLoc:                srcLoc,
		dstLoc:                dstLoc,
		status:                common.TripStatusDriverOnTheWay,
		tripID:                tripID,
		price:                 price,
		pricingStrategy:       pricingStrategy,
		driverMatchingStrategy: driverMatchingStrategy,
	}
}

// GetTripID returns the trip ID
func (t *Trip) GetTripID() int {
	return t.tripID
}

// DisplayTripDetails displays the trip details
func (t *Trip) DisplayTripDetails() {
	fmt.Println()
	fmt.Printf("Trip id - %d\n", t.tripID)
	fmt.Printf("Rider - %s\n", t.rider.GetRiderName())
	fmt.Printf("Driver - %s\n", t.driver.GetDriverName())
	fmt.Printf("Price - %.2f\n", t.price)
	fmt.Printf("Locations - %.2f,%.2f and %.2f,%.2f\n",
		t.srcLoc.Latitude, t.srcLoc.Longitude,
		t.dstLoc.Latitude, t.dstLoc.Longitude)
}

