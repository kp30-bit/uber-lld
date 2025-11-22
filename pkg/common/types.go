package common

import "sync"

// Rating represents the rating level
type Rating int

const (
	Unassigned Rating = iota
	OneStar
	TwoStars
	ThreeStars
	FourStars
	FiveStars
)

// TripStatus represents the status of a trip
type TripStatus int

const (
	TripStatusUnassigned TripStatus = iota
	TripStatusDriverOnTheWay
	TripStatusDriverArrived
	TripStatusStarted
	TripStatusPaused
	TripStatusCancelled
	TripStatusEnded
)

// RatingToString converts a Rating to its string representation
func RatingToString(rating Rating) string {
	switch rating {
	case OneStar:
		return "one star"
	case TwoStars:
		return "two stars"
	case ThreeStars:
		return "three stars"
	case FourStars:
		return "four stars"
	case FiveStars:
		return "five stars"
	default:
		return "invalid rating"
	}
}

// IsHighRating checks if a rating is high (4 or 5 stars)
func IsHighRating(rating Rating) bool {
	return rating == FourStars || rating == FiveStars
}

var (
	nextTripID  int = 0
	tripIDMutex sync.Mutex
)

// GetNextTripID returns the next trip ID in a thread-safe manner
func GetNextTripID() int {
	tripIDMutex.Lock()
	defer tripIDMutex.Unlock()
	nextTripID++
	return nextTripID
}
