package domain

import (
	"log"

	"github.com/kamalpratik/Uber-Ola-Low-Level-Design/pkg/common"
)

type UserType int

const (
	UserTypeRider UserType = iota
	UserTypeDriver
)

func CreateRider(userType UserType, name string, rating common.Rating) *Rider {
	user, err := ActorFactory(userType, name, rating)
	if err != nil {
		log.Fatalf("failed to create rider: %v", err)
	}
	return user.(*Rider)
}

// Helper function to create and assert driver
func CreateDriver(userType UserType, name string, rating common.Rating) *Driver {
	user, err := ActorFactory(userType, name, rating)
	if err != nil {
		log.Fatalf("failed to create driver: %v", err)
	}
	return user.(*Driver)
}
