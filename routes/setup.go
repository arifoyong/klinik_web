package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine, port string) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Main page"})
	})

	// NotaRoutes(router)
	PatientRoutes(router)
	VisitRoutes(router)
	DrugRoutes(router)

	router.Run("localhost:" + port)
}
