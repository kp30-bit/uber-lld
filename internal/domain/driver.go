package domain

import "github.com/kamalpratik/Uber-Ola-Low-Level-Design/pkg/common"

// Driver represents a driver in the system
type Driver struct {
	name   string
	avail  bool
	rating common.Rating
}

// NewDriver creates a new Driver instance
func NewDriver(name string, rating common.Rating) *Driver {
	return &Driver{
		name:   name,
		avail:  false,
		rating: rating,
	}
}

// UpdateAvail updates the driver's availability
func (d *Driver) UpdateAvail(avail bool) {
	d.avail = avail
}

// GetDriverName returns the driver's name
func (d *Driver) GetDriverName() string {
	return d.name
}

// GetRating returns the driver's rating
func (d *Driver) GetRating() common.Rating {
	return d.rating
}

