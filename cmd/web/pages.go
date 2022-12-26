package main

import (
	"coursework/pkg/handlers/flight"
	"coursework/pkg/sevices/token"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func homePage(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		title := "AviaSales"
		getUser, err := token.ParseToken(tokenStr)

		if err != nil {
			tokenStr = ""
			title = "Авторизация"

			c.HTML(http.StatusOK, "index.html", gin.H{
				"title": title,
				"id":    getUser.Id,
				"role":  getUser.Role,
			})
		} else {
			flightStruct := flight.GetAllFlights()

			c.HTML(http.StatusOK, "index.html", gin.H{
				"title":   title,
				"id":      getUser.Id,
				"role":    getUser.Role,
				"flight":  flightStruct,
				"message": "Все рейсы",
			})
		}
	})

	router.GET("/search", func(c *gin.Context) {
		getUser, err := token.ParseToken(tokenStr)

		if err != nil && getUser.Role == 0 {
			panic("Don't have rools to do this")
		}

		destination := c.Query("destination")
		typeClass := c.Query("type_class")

		flightStruct, message := flight.GetAllFlightsByUser(destination, typeClass)

		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":   "AviaSales",
			"id":      getUser.Id,
			"role":    getUser.Role,
			"flight":  flightStruct,
			"message": message,
		})
	})
}

func registrationPage(router *gin.Engine) {
	router.GET("/registration", func(c *gin.Context) {

		c.HTML(http.StatusOK, "registration.html", gin.H{
			"title": "Регистрация",
		})
	})
}

func flightPage(router *gin.Engine) {
	router.GET("/flight", func(c *gin.Context) {
		if len(tokenStr) == 0 {
			panic("Don't have rools to do this")
		}

		c.HTML(http.StatusOK, "flight_form.html", gin.H{
			"title":  "Создать рейс",
			"method": "add",
		})
	})

	router.GET("/flight/:id", func(c *gin.Context) {
		if len(tokenStr) == 0 {
			panic("Don't have rools to do this")
		}

		id, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			panic("Invalid id")
		}

		flightStruct := flight.GetFlightById(uint(id))

		c.HTML(http.StatusOK, "flight_form.html", gin.H{
			"title":  "Создать рейс",
			"flight": flightStruct,
			"method": "update",
		})
	})
}

func ticketPage(router *gin.Engine) {
	router.GET("/ticket/:id", func(c *gin.Context) {
		if len(tokenStr) == 0 {
			panic("Don't have rools to do this")
		}

		id, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			panic("Invalid id")
		}

		flightStruct := flight.GetFlightById(uint(id))

		c.HTML(http.StatusOK, "ticket_form.html", gin.H{
			"title":  "Купить билет",
			"flight": flightStruct,
		})
	})
}
