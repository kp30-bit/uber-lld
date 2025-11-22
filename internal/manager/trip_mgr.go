package manager

import (
	"sync"
	"github.com/kamalpratik/Uber-Ola-Low-Level-Design/internal/domain"
)

// TripMgr manages trips in the system (Singleton pattern)
type TripMgr struct {
	riderMgr          *RiderMgr
	driverMgr         *DriverMgr
	tripsMetaDataInfo map[int]*domain.TripMetaData
	tripsInfo         map[int]*domain.Trip
	mu                sync.RWMutex
}

var (
	tripMgrInstance *TripMgr
	tripMgrOnce     sync.Once
)

// GetTripMgr returns the singleton instance of TripMgr
func GetTripMgr() *TripMgr {
	tripMgrOnce.Do(func() {
		tripMgrInstance = &TripMgr{
			riderMgr:          GetRiderMgr(),
			driverMgr:         GetDriverMgr(),
			tripsMetaDataInfo: make(map[int]*domain.TripMetaData),
			tripsInfo:         make(map[int]*domain.Trip),
		}
	})
	return tripMgrInstance
}

// CreateTrip creates a new trip
func (t *TripMgr) CreateTrip(rider *domain.Rider, srcLoc, dstLoc *domain.Location) {
	metaData := domain.NewTripMetaData(srcLoc, dstLoc, rider.GetRating())
	strategyMgr := GetStrategyMgr()
	
	pricingStrategy := strategyMgr.DeterminePricingStrategy(metaData)
	driverMatchingStrategy := strategyMgr.DetermineMatchingStrategy(metaData)
	
	driver := driverMatchingStrategy.MatchDriver(metaData)
	if driver == nil {
		return // No driver available
	}
	
	tripPrice := pricingStrategy.CalculatePrice(metaData)
	
	trip := domain.NewTrip(rider, driver, srcLoc, dstLoc, tripPrice, pricingStrategy, driverMatchingStrategy)
	tripID := trip.GetTripID()
	
	t.mu.Lock()
	defer t.mu.Unlock()
	t.tripsInfo[tripID] = trip
	t.tripsMetaDataInfo[tripID] = metaData
}

// GetTripsMap returns the trips map
func (t *TripMgr) GetTripsMap() map[int]*domain.Trip {
	t.mu.RLock()
	defer t.mu.RUnlock()
	// Return a copy to prevent external modifications
	result := make(map[int]*domain.Trip)
	for k, v := range t.tripsInfo {
		result[k] = v
	}
	return result
}

