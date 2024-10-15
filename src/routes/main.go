package routes

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

var router = gin.Default()

const logLabel = "[GIN-Router]"

func InitRoutes() {
	v1 := router.Group("/v1")
	defineUserRoutes(v1);
}

func Run() {
	InitRoutes()

	port := os.Getenv("PORT")

	if (port == "") {
		port = "8080"
	}

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	fmt.Printf("%s Starting Gin router on %s\n", logLabel, port)
	router.Run(":" + port)
}