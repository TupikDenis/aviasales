package flight

import (
	"coursework/pkg/handlers"
	"coursework/pkg/models/databaseModels/flight"
)

func CreateFlight(flightNumber uint, typeFlight string, destination string, date string,
	startTime string, endTime string, capacity uint, businessCapacity uint, businessCost uint,
	economCapacity uint, economCost uint) {
	planeAdd := flight.Flight{
		FlightNumber:     flightNumber,
		TypeFlight:       typeFlight,
		Destination:      destination,
		Date:             date,
		StartTime:        startTime,
		EndTime:          endTime,
		Capacity:         capacity,
		BusinessCapacity: businessCapacity,
		BusinessCost:     businessCost,
		EconomCapacity:   economCapacity,
		EconomCost:       economCost,
	}

	db := handlers.Database()

	err := db.AutoMigrate(&flight.Flight{})
	if err != nil {
		panic("Error: I can't add new flight")
	}

	db.Save(&planeAdd)
}
