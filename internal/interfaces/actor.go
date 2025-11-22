package interfaces

import "github.com/kamalpratik/Uber-Ola-Low-Level-Design/pkg/common"

type Actor interface {
	GetName() string
	GetRating() common.Rating
}
