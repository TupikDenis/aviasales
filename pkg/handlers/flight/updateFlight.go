package flight

import (
	"coursework/pkg/handlers"
	"coursework/pkg/models/databaseModels/flight"
)

func UpdateFlight(id uint, flightNumber uint, typeFlight string, destination string, date string,
	startTime string, endTime string, capacity uint, businessCapacity uint, businessCost uint,
	economCapacity uint, economCost uint) {
	var flightStruct flight.Flight
	db := handlers.Database()

	db.First(&flightStruct, id)

	flightStruct.FlightNumber = flightNumber
	flightStruct.TypeFlight = typeFlight
	flightStruct.Destination = destination
	flightStruct.Date = date
	flightStruct.StartTime = startTime
	flightStruct.EndTime = endTime
	flightStruct.Capacity = capacity
	flightStruct.StartTime = startTime
	flightStruct.EndTime = endTime
	flightStruct.BusinessCapacity = businessCapacity
	flightStruct.BusinessCost = businessCost
	flightStruct.EconomCapacity = economCapacity
	flightStruct.EconomCost = economCost

	db.Save(&flightStruct)
}

func UpdateTicket(id uint, businessCapacity uint, economCapacity uint) {
	var flightStruct flight.Flight
	db := handlers.Database()

	db.First(&flightStruct, id)

	flightStruct.BusinessCapacity -= businessCapacity
	flightStruct.EconomCapacity -= economCapacity

	db.Save(&flightStruct)
}
