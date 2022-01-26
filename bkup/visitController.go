// package controllers

// import (
// 	"net/http"

// 	"github.com/arifoyong/klinik/models"
// 	"github.com/gin-gonic/gin"
// )

// type CreateVisitVal struct {
// 	// Date            time.Time `json:"date"`
// 	Patient_id      uint   `json:"patient_id" validate:"required"`
// 	Problems        string `json:"problems"`
// 	Diagnosis       string `json:"diagnosis"`
// 	Prescription_id uint   `json:"prescription_id"`
// }

// // ListAllVisits get all visits
// // and return as JSON
// func GetAllVisits(c *gin.Context) {
// 	var visits []models.Visit

// 	if err := models.DB.Find(&visits).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": visits})
// }

// // GetVisitById get visit by id specified in parameter
// // and return it as JSON
// func GetVisitById(c *gin.Context) {
// 	var visit models.Visit

// 	if err := models.DB.Where("id = ?", c.Param("id")).First(&visit).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": visit})
// }

// // AddVisit add visit with json data provided in request body
// // and return the status as JSON
// func AddVisit(c *gin.Context) {
// 	var input CreateVisitVal
// 	var visit models.Visit

// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	visit.Patient_id = input.Patient_id
// 	visit.Problems = input.Problems
// 	visit.Diagnosis = input.Diagnosis
// 	visit.Prescription_id = input.Prescription_id

// 	if err := models.DB.Create(&visit).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": "add success"})
// }

// func EditVisit(c *gin.Context)   {}
// func DeleteVisit(c *gin.Context) {}
