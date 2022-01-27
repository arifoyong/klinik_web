package routes

import (
	"github.com/arifoyong/klinik/controllers"
	"github.com/gin-gonic/gin"
)

func DrugRoutes(router *gin.Engine) {
	router.GET("/drugs", controllers.GetAllDrugs)
	router.GET("/drugs/:id", controllers.GetDrugById)
	router.POST("/drugs", controllers.AddDrug)
	router.PATCH("/drugs/:id", controllers.EditDrug)
	router.DELETE("/drugs/:id", controllers.DeleteDrug)
}
