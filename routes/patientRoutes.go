package routes

import (
	"github.com/arifoyong/klinik/controllers"
	"github.com/gin-gonic/gin"
)

func PatientRoutes(router *gin.Engine) {
	router.GET("/patients", controllers.GetPatients)
	router.GET("/patients/name/:name", controllers.GetPatientByName)
	router.GET("/patients/id/:id", controllers.GetPatientById)
	// router.POST("/patients", controllers.AddPatient)
	// router.PATCH("/patients/:id", controllers.EditPatient)
	// router.DELETE("/patients/:id", controllers.DeletePatient)
}
