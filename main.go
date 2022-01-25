package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/arifoyong/klinik/controllers"
	"github.com/arifoyong/klinik/models"
)

func main() {
	router := gin.Default()
	models.ConnectDB()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "Main page"})
	})
	router.GET("/notas", controllers.ListNotas)
	router.POST("/notas", controllers.CreateNota)
	router.GET("/notas/:id", controllers.GetNotaById)
	router.PATCH("/notas/:id", controllers.UpdateNota)
	router.DELETE(("/notas/:id"), controllers.DeleteNota)

	router.GET("/patients", controllers.GetPatients)
	router.GET("/patients/name/:name", controllers.GetPatientByName)
	router.GET("/patients/id/:id", controllers.GetPatientById)
	router.POST("/patients", controllers.AddPatient)
	router.PATCH("/patients/:id", controllers.EditPatient)
	router.DELETE("/patients/:id", controllers.DeletePatient)

	router.Run("localhost:8000")
}
