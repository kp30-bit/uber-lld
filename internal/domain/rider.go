package domain

import "github.com/kamalpratik/Uber-Ola-Low-Level-Design/pkg/common"

// Rider represents a rider in the system
type Rider struct {
	name   string
	rating common.Rating
}

// NewRider creates a new Rider instance
func NewRider(name string, rating common.Rating) *Rider {
	return &Rider{
		name:   name,
		rating: rating,
	}
}

// GetRiderName returns the rider's name
func (r *Rider) GetName() string {
	return r.name
}

// GetRating returns the rider's rating
func (r *Rider) GetRating() common.Rating {
	return r.rating
}
