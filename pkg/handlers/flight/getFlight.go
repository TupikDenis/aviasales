package flight

import (
	"coursework/pkg/handlers"
	"coursework/pkg/models/databaseModels/flight"
	"strconv"
)

func GetFlightById(id uint) flight.TransformedFlight {
	var flights []flight.Flight
	var flightById flight.TransformedFlight

	db := handlers.Database()
	db.Where("id = ?", strconv.Itoa(int(id))).Find(&flights)

	flightById = flight.TransformedFlight{
		Id:               flights[0].ID,
		FlightNumber:     flights[0].FlightNumber,
		TypeFlight:       flights[0].TypeFlight,
		Destination:      flights[0].Destination,
		Date:             flights[0].Date,
		StartTime:        flights[0].StartTime,
		EndTime:          flights[0].EndTime,
		Capacity:         flights[0].Capacity,
		BusinessCapacity: flights[0].BusinessCapacity,
		BusinessCost:     flights[0].BusinessCost,
		EconomCapacity:   flights[0].EconomCapacity,
		EconomCost:       flights[0].EconomCost,
	}

	return flightById
}
