package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/arifoyong/klinik/models"
	"github.com/arifoyong/klinik/routes"
)

func main() {
	port := "8000"
	router := gin.Default()

	// SETUP CORS
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:3000"}
	corsConfig.AllowCredentials = true
	corsConfig.AddAllowMethods("OPTIONS") // OPTIONS method for ReactJS
	router.Use(cors.New(corsConfig))      // Register the middleware

	// Initialize DB
	models.ConnectDatabase()

	// Setup router
	routes.Setup(router, port)
}
