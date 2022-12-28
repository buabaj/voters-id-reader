package main

import (
	"github.com/buabaj/voters-id-reader/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", controllers.Root)
	r.POST("/read", controllers.ReadID)

	r.Use(cors.Default())
	r.Run() // listen and serve on 8080
}
