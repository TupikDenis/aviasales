package main

var tokenStr string
var port string = ":8080"

func handle() {
	router := createRouter()

	homePage(router)
	registrationPage(router)
	flightPage(router)
	ticketPage(router)
	userPage(router)

	userRouterGroup(router)
	flightsRouterGroup(router)
	ticketsRouterGroup(router)

	startServer(router, port)
}
