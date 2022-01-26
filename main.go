package main

import (
	"github.com/gin-gonic/gin"

	"github.com/arifoyong/klinik/routes"
)

func main() {
	port := "8000"
	router := gin.Default()
	// models.SetupDB()

	routes.Setup(router, port)
}
