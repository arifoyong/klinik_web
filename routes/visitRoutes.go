package routes

import (
	"github.com/arifoyong/klinik/controllers"
	"github.com/gin-gonic/gin"
)

func VisitRoutes(router *gin.Engine) {
	router.GET("/visits", controllers.GetAllVisits)
	router.GET("/visits/id/:id", controllers.GetVisitById)
	router.GET("/visits/status/:status", controllers.GetVisitByStatus)
	router.POST("/visits", controllers.AddVisit)
	router.DELETE("/visits/:id", controllers.DeleteVisit)
}
