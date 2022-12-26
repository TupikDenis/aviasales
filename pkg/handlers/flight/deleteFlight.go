package flight

import (
	"coursework/pkg/handlers"
	"coursework/pkg/models/databaseModels/flight"
)

func DeleteFlight(id uint) {
	var flightStruct flight.Flight
	db := handlers.Database()
	db.First(&flightStruct, id)
	db.Delete(&flightStruct)
}
