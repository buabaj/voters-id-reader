package main

import (
	"github.com/buabaj/voters-id-reader/controllers"
	"github.com/buabaj/voters-id-reader/utilities"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(utilities.CORS())
	r.GET("/", controllers.Root)
	r.POST("/read", controllers.ReadID)

	r.Run() // listen and serve on 8080
}
