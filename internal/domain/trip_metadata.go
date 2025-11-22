package domain

import "github.com/kamalpratik/Uber-Ola-Low-Level-Design/pkg/common"

// TripMetaData contains metadata for a trip that will be used by strategies
// Pricing strategy and Driver Matching strategy will need this data
// Even if more data is needed, only this class needs to be updated
type TripMetaData struct {
	srcLoc      *Location
	dstLoc      *Location
	riderRating common.Rating
	driverRating common.Rating
}

// NewTripMetaData creates a new TripMetaData instance
func NewTripMetaData(srcLoc, dstLoc *Location, riderRating common.Rating) *TripMetaData {
	return &TripMetaData{
		srcLoc:       srcLoc,
		dstLoc:       dstLoc,
		riderRating:  riderRating,
		driverRating: common.Unassigned,
	}
}

// GetRiderRating returns the rider's rating
func (t *TripMetaData) GetRiderRating() common.Rating {
	return t.riderRating
}

// GetDriverRating returns the driver's rating
func (t *TripMetaData) GetDriverRating() common.Rating {
	return t.driverRating
}

// SetDriverRating sets the driver's rating
func (t *TripMetaData) SetDriverRating(driverRating common.Rating) {
	t.driverRating = driverRating
}

