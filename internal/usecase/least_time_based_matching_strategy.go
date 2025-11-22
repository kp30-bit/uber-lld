package strategy

import (
	"fmt"

	"github.com/kamalpratik/Uber-Ola-Low-Level-Design/internal/domain"
	"github.com/kamalpratik/Uber-Ola-Low-Level-Design/pkg/interfaces"
)

// LeastTimeBasedMatchingStrategy implements driver matching based on least time
type LeastTimeBasedMatchingStrategy struct {
	driverProvider interfaces.DriverProvider
}

// NewLeastTimeBasedMatchingStrategy creates a new LeastTimeBasedMatchingStrategy instance
func NewLeastTimeBasedMatchingStrategy(driverProvider interfaces.DriverProvider) interfaces.DriverMatchingStrategy {
	return &LeastTimeBasedMatchingStrategy{
		driverProvider: driverProvider,
	}
}

// MatchDriver matches a driver based on least time
func (l *LeastTimeBasedMatchingStrategy) MatchDriver(tripMetaData *domain.TripMetaData) *domain.Driver {
	driversMap := l.driverProvider.GetDriversMap()

	if len(driversMap) == 0 {
		fmt.Println("No drivers! What service is this huh?")
		return nil
	}

	fmt.Println("Using quadtree to see nearest cabs, using driver manager to get details of drivers and send notifications")

	// Get the first driver from the map (here we can call quadtree algo to get nearest cabs)
	var driver *domain.Driver
	for _, d := range driversMap {
		driver = d
		break
	}

	fmt.Printf("Setting %s as driver\n", driver.GetName())
	tripMetaData.SetDriverRating(driver.GetRating())
	return driver
}
