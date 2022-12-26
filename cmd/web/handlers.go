package main

// run project  go run ./cmd/web/.

var tokenStr string
var port string = ":8080"

func handle() {
	router := createRouter()

	homePage(router)
	registrationPage(router)
	flightPage(router)
	ticketPage(router)

	userRouterGroup(router)
	flightsRouterGroup(router)
	ticketsRouterGroup(router)

	startServer(router, port)
}
