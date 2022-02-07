package controllers

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/arifoyong/klinik/models"
	"github.com/gin-gonic/gin"
)

type GetVisitQry struct {
	ID              uint      `json:"id"`
	Date            time.Time `json:"date"`
	Patient_id      uint      `json:"patient_id"`
	Firstname       *string   `json:"firstname"`
	Lastname        *string   `json:"lastname"`
	IC              *string   `json:"ic" `
	Problems        *string   `json:"problems"`
	Diagnosis       *string   `json:"diagnosis"`
	Prescription_id uint      `json:"prescription_id"`
}

// ListAllVisits get all visits
// and return as JSON
func GetAllVisits(c *gin.Context) {
	sqlStatement := `SELECT visits.id, visits.date, visits.patient_id, visits.problems, visits.diagnosis, visits.prescription_id,
										patients.firstname, patients.lastname, patients.ic  
									FROM visits INNER JOIN	patients ON (visits.patient_id = patients.id) `
	rows, err := models.DB.Query(sqlStatement)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var visits []GetVisitQry
	for rows.Next() {
		var visit GetVisitQry
		err := rows.Scan(&visit.ID, &visit.Date, &visit.Patient_id, &visit.Problems, &visit.Diagnosis, &visit.Prescription_id, &visit.Firstname, &visit.Lastname, &visit.IC)
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
	var visit GetVisitQry

	sqlStatement := `SELECT visits.id, visits.date, visits.patient_id, visits.problems, visits.diagnosis, visits.prescription_id,
													patients.firstname, patients.lastname, patients.ic  
									FROM visits INNER JOIN	patients ON (visits.patient_id = patients.id) 
									WHERE visits.id = $1`
	err := models.DB.QueryRow(sqlStatement, c.Param("id")).Scan(&visit.ID, &visit.Date, &visit.Patient_id,
		&visit.Problems, &visit.Diagnosis, &visit.Prescription_id, &visit.Firstname, &visit.Lastname, &visit.IC)
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
	var input models.Visit
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	sqlStatement := `INSERT INTO visits(date, patient_id, problems, diagnosis, prescription_id)
	VALUES($1, $2, $3, $4, $5)`

	_, err := models.DB.Exec(sqlStatement, input.Date, input.Patient_id, input.Problems, input.Diagnosis, input.Prescription_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "add success"})
}

// func EditVisit(c *gin.Context)   {}
func DeleteVisit(c *gin.Context) {
	sqlStatement := `DELETE FROM visits WHERE id = $1`
	_, err := models.DB.Exec(sqlStatement, c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "delete success"})
}
