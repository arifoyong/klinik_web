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

	router.Run("localhost:8000")
}
