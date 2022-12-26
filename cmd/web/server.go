package main

import "github.com/gin-gonic/gin"

func startServer(router *gin.Engine, port string) {
	err := router.Run(port)

	if err != nil {
		panic(err)
	}
}
