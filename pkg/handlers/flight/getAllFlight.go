package flight

import (
	"coursework/pkg/handlers"
	"coursework/pkg/models/databaseModels/flight"
	"strconv"
)

func GetAllFlights() []flight.TransformedFlight {
	var flights []flight.Flight

	db := handlers.Database()
	db.Order("date asc").Find(&flights)

	return getFlights(flights)
}

func GetAllFlightsByUser(destination string, typeClass string) ([]flight.TransformedFlight, string) {
	var flights []flight.Flight

	db := handlers.Database()

	switch typeClass {
	case "business":
		db.Order("date asc").Where("destination = ?", destination).Where("business_capacity > 0").Find(&flights)
	case "econom":
		db.Order("date asc").Where("destination = ?", destination).Where("econom_capacity > 0").Find(&flights)
	}

	if len(getFlights(flights)) == 0 && typeClass == "business" {
		db.Order("date asc").Where("destination = ?", destination).Where("econom_capacity > 0").Find(&flights)

		if len(getFlights(flights)) == 0 {
			return getFlights(flights), "К сожалению, мы ничего не нашли"
		}

		var ticket uint
		for _, item := range flights {
			ticket += item.EconomCapacity
		}

		return getFlights(flights), "К сожалению, мы не нашли для вас билеты бизнес-класса, но на найденные рейсы есть билеты эконом-класса " +
			"в количестве " + strconv.Itoa(int(ticket))
	} else if len(getFlights(flights)) == 0 && typeClass == "econom" {
		db.Order("date asc").Where("destination = ?", destination).Where("business_capacity > 0").Find(&flights)
		if len(getFlights(flights)) == 0 {
			return getFlights(flights), "К сожалению, мы ничего не нашли"
		}

		var ticket uint
		for _, item := range flights {
			ticket += item.BusinessCapacity
		}

		return getFlights(flights), "К сожалению, мы не нашли для вас билеты эконом-класса, но на найденные рейсы есть билеты бизнес-класса " +
			"в количестве " + strconv.Itoa(int(ticket))
	} else {
		return getFlights(flights), "По вашему поиску найдено следующее:"
	}
}

func getFlights(flights []flight.Flight) []flight.TransformedFlight {
	var _flights []flight.TransformedFlight

	for _, item := range flights {
		_flights = append(
			_flights, flight.TransformedFlight{
				Id:               item.ID,
				FlightNumber:     item.FlightNumber,
				TypeFlight:       item.TypeFlight,
				Destination:      item.Destination,
				Date:             item.Date,
				StartTime:        item.StartTime,
				EndTime:          item.EndTime,
				Capacity:         item.Capacity,
				BusinessCapacity: item.BusinessCapacity,
				BusinessCost:     item.BusinessCost,
				EconomCapacity:   item.EconomCapacity,
				EconomCost:       item.EconomCost,
			})
	}

	return _flights
}
