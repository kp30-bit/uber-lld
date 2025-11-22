package main

import (
	"fmt"
	"github.com/kamalpratik/Uber-Ola-Low-Level-Design/internal/domain"
	"github.com/kamalpratik/Uber-Ola-Low-Level-Design/internal/manager"
	"github.com/kamalpratik/Uber-Ola-Low-Level-Design/pkg/common"
)

func main() {
	//---------------Creating Riders and Drivers--------------------------------
	keertiRider := domain.NewRider("Keerti", common.FiveStars)
	gauravRider := domain.NewRider("Gaurav", common.FiveStars)
	riderMgr := manager.GetRiderMgr()
	riderMgr.AddRider("keerti", keertiRider)
	riderMgr.AddRider("gaurav", gauravRider)

	yogitaDriver := domain.NewDriver("Yogita", common.ThreeStars)
	riddhiDriver := domain.NewDriver("Riddhi", common.FourStars)
	driverMgr := manager.GetDriverMgr()
	driverMgr.AddDriver("yogita", yogitaDriver)
	driverMgr.AddDriver("riddhi", riddhiDriver)

	//These details in turn will be stored in database
	//-------------------------------------------------------------------------

	tripMgr := manager.GetTripMgr()

	fmt.Println()
	fmt.Println("Creating Trip for Keerti from location (10,10) to (30,30)")
	tripMgr.CreateTrip(keertiRider, domain.NewLocation(10, 10), domain.NewLocation(30, 30))

	fmt.Println()
	fmt.Println("Creating Trip for Gaurav from location (200,200) to (500,500)")
	tripMgr.CreateTrip(gauravRider, domain.NewLocation(200, 200), domain.NewLocation(500, 500))

	//-------------------Display all the trips created--------------------------
	tripsMap := tripMgr.GetTripsMap()
	for _, trip := range tripsMap {
		trip.DisplayTripDetails()
	}
}

