package domain

import (
	"fmt"

	"github.com/kamalpratik/Uber-Ola-Low-Level-Design/internal/interfaces"
	"github.com/kamalpratik/Uber-Ola-Low-Level-Design/pkg/common"
)

func ActorFactory(userType UserType, name string, rating common.Rating) (interfaces.Actor, error) {
	switch userType {
	case UserTypeRider:
		return NewRider(name, rating), nil
	case UserTypeDriver:
		return NewDriver(name, rating), nil
	default:
		return nil, fmt.Errorf("unknown user type: %v", userType)
	}
}
