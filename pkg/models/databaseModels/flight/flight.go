package flight

import (
	"gorm.io/gorm"
)

type Flight struct {
	gorm.Model
	FlightNumber     uint   `json:"flight_number" gorm:"uniqueIndex"`
	TypeFlight       string `json:"type_flight"`
	Destination      string `json:"destination"`
	Date             string `json:"date"`
	StartTime        string `json:"start_time"`
	EndTime          string `json:"end_time"`
	Capacity         uint   `json:"capacity"`
	BusinessCapacity uint   `json:"business_capacity"`
	BusinessCost     uint   `json:"business_cost"`
	EconomCapacity   uint   `json:"econom_capacity"`
	EconomCost       uint   `json:"econom_cost"`
}

type TransformedFlight struct {
	Id               uint   `json:"id"`
	FlightNumber     uint   `json:"flight_number"`
	TypeFlight       string `json:"type_flight"`
	Destination      string `json:"destination"`
	Date             string `json:"date"`
	StartTime        string `json:"start_time"`
	EndTime          string `json:"end_time"`
	Capacity         uint   `json:"capacity"`
	BusinessCapacity uint   `json:"business_capacity"`
	BusinessCost     uint   `json:"business_cost"`
	EconomCapacity   uint   `json:"econom_capacity"`
	EconomCost       uint   `json:"econom_cost"`
}
