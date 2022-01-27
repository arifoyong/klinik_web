package controllers

import (
	"database/sql"
	"net/http"

	"github.com/arifoyong/klinik/models"
	"github.com/gin-gonic/gin"
)

// ListAllVisits get all visits
// and return as JSON
func GetAllVisits(c *gin.Context) {
	db := models.SetupDB()

	sqlStatement := `SELECT * FROM visits`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var visits []models.Visit
	for rows.Next() {
		var visit models.Visit
		err := rows.Scan(&visit.ID, &visit.Date, &visit.Patient_id, &visit.Problems, &visit.Diagnosis, &visit.Prescription_id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		visits = append(visits, visit)
	}

	c.JSON(http.StatusOK, gin.H{"data": visits})
}

// GetVisitById get visit by id specified in parameter
// and return it as JSON
func GetVisitById(c *gin.Context) {
	db := models.SetupDB()
	var visit models.Visit

	sqlStatement := `SELECT * FROM visits WHERE id = $1`
	err := db.QueryRow(sqlStatement, c.Param("id")).Scan(&visit.ID, &visit.Date, &visit.Patient_id, &visit.Problems, &visit.Diagnosis, &visit.Prescription_id)
	switch err {
	case sql.ErrNoRows:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
	case nil:
		c.JSON(http.StatusOK, visit)
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

// AddVisit add visit with json data provided in request body
// and return the status as JSON
func AddVisit(c *gin.Context) {
	db := models.SetupDB()

	var input models.Visit
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	sqlStatement := `INSERT INTO visits(date, patient_id, problems, diagnosis, prescription_id)
	VALUES($1, $2, $3, $4, $5)`

	_, err := db.Exec(sqlStatement, input.Date, input.Patient_id, input.Problems, input.Diagnosis, input.Prescription_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "add success"})
}

// func EditVisit(c *gin.Context)   {}
func DeleteVisit(c *gin.Context) {
	db := models.SetupDB()

	sqlStatement := `DELETE FROM visits WHERE id = $1`
	_, err := db.Exec(sqlStatement, c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "delete success"})
}
