package main

import (
	"coursework/pkg/handlers/flight"
	"coursework/pkg/handlers/user"
	"coursework/pkg/sevices/token"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func createRouter() *gin.Engine {
	router := gin.Default()
	router.Static("/static", "./ui/static")
	router.LoadHTMLGlob("ui/html/*")

	return router
}

func userRouterGroup(router *gin.Engine) {
	enterExit := router.Group("")
	{
		enterExit.POST("/authentication", func(c *gin.Context) {
			login := c.PostForm("login")
			password := c.PostForm("password")

			tokenStr = user.Login(login, password)
			c.Redirect(http.StatusFound, "/")
		})

		enterExit.GET("/logout", func(c *gin.Context) {
			tokenStr = ""
			c.Redirect(http.StatusFound, "/")
		})
	}

	users := router.Group("/api/users")
	{
		users.POST("/", func(c *gin.Context) {
			username := c.PostForm("login")
			password := c.PostForm("password")

			isCreateUser := user.CreateUser(username, password)

			if isCreateUser {
				c.Redirect(http.StatusFound, "/")
			} else {
				c.Redirect(http.StatusFound, "/registration")
			}
		})
	}
}

func flightsRouterGroup(router *gin.Engine) {
	flights := router.Group("/api/flights")
	{
		flights.POST("/", func(c *gin.Context) {
			getUser, err := token.ParseToken(tokenStr)

			if err != nil && getUser.Role == 0 {
				panic("Don't have rools to do this")
			}

			flightNumber, typeFlight, destination, date, startTime,
				endTime, capacity, businessCapacity, businessCost, economCapacity, economCost := getFieldsForUpdateAndCreate(c)

			flight.CreateFlight(flightNumber, typeFlight, destination, date, startTime, endTime,
				capacity, businessCapacity, businessCost, economCapacity, economCost)

			c.Redirect(http.StatusFound, "/")
		})

		flights.PUT("/:id", func(c *gin.Context) {
			getUser, err := token.ParseToken(tokenStr)

			if err != nil && getUser.Role == 0 {
				panic("Don't have rools to do this")
			}

			id, err := strconv.ParseUint(c.Param("id"), 10, 64)
			if err != nil {
				panic("Invalid id")
			}

			flightNumber, typeFlight, destination, date, startTime,
				endTime, capacity, businessCapacity, businessCost, economCapacity, economCost := getFieldsForUpdateAndCreate(c)

			flight.UpdateFlight(uint(id), flightNumber, typeFlight, destination, date, startTime, endTime,
				capacity, businessCapacity, businessCost, economCapacity, economCost)
		})

		flights.DELETE("/:id", func(c *gin.Context) {
			getUser, err := token.ParseToken(tokenStr)

			if err != nil && getUser.Role == 0 {
				panic("Don't have rools to do this")
			}

			id, err := strconv.ParseUint(c.Param("id"), 10, 64)
			if err != nil {
				panic("Invalid id")
			}

			flight.DeleteFlight(uint(id))
		})
	}

}

func getFieldsForUpdateAndCreate(c *gin.Context) (uint, string, string, string, string, string, uint, uint, uint, uint, uint) {
	flightNumber, err := strconv.ParseUint(c.PostForm("number"), 10, 64)
	if err != nil {
		panic("Bad flight number")
	}

	typeflight := c.PostForm("type")
	switch typeflight {
	case "citizen":
		typeflight = "Пассажирский"
	case "war":
		typeflight = "Военный"
	}

	destination := c.PostForm("destination")
	date := c.PostForm("date")
	startTime := c.PostForm("start_time")
	endTime := c.PostForm("end_time")

	capacity, err := strconv.ParseUint(c.PostForm("capacity"), 10, 64)
	if err != nil {
		panic("Bad capacity")
	}

	businessCapacity, err := strconv.ParseUint(c.PostForm("business_capacity"), 10, 64)
	if err != nil {
		panic("Bad capacity (business)")
	}

	businessCost, err := strconv.ParseUint(c.PostForm("business_cost"), 10, 64)
	if err != nil {
		panic("Bad cost (business)")
	}

	economCapacity, err := strconv.ParseUint(c.PostForm("econom_capacity"), 10, 64)
	if err != nil {
		panic("Bad capacity (econom)")
	}

	economCost, err := strconv.ParseUint(c.PostForm("econom_cost"), 10, 64)
	if err != nil {
		panic("Bad cost (econom)")
	}

	return uint(flightNumber), typeflight, destination, date, startTime, endTime,
		uint(capacity), uint(businessCapacity), uint(businessCost), uint(economCapacity), uint(economCost)
}

func ticketsRouterGroup(router *gin.Engine) {
	router.PATCH("/api/tickets/:id", func(c *gin.Context) {
		getUser, err := token.ParseToken(tokenStr)

		if err != nil && getUser.Role == 0 {
			panic("Don't have rools to do this")
		}

		id, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			panic("Invalid id")
		}

		businessCapacity, err := strconv.ParseUint(c.PostForm("business_cost"), 10, 64)
		if err != nil {
			panic("Bad capacity (business)")
		}

		economCapacity, err := strconv.ParseUint(c.PostForm("econom_cost"), 10, 64)
		if err != nil {
			panic("Bad capacity (econom)")
		}

		flight.UpdateTicket(uint(id), uint(businessCapacity), uint(economCapacity))
	})
}
