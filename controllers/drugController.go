package controllers

import (
	"database/sql"
	"net/http"

	"github.com/arifoyong/klinik/models"
	"github.com/gin-gonic/gin"
)

// AddDrug add drugs based on JSON information provided
// in request body and return a JSON response
func AddDrug(c *gin.Context) {
	db := models.SetupDB()

	var input models.Drug
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	sqlStatement := `INSERT INTO drugs(name, unit_price) VALUES($1, $2)`
	_, err := db.Exec(sqlStatement, input.Name, input.Unit_price)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Add success"})
}

// GetAllDrugs get all drugs and
// return as JSON
func GetAllDrugs(c *gin.Context) {
	db := models.SetupDB()

	sqlStatement := `SELECT * FROM drugs`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var drugs []models.Drug
	for rows.Next() {
		var drug models.Drug
		if err := rows.Scan(&drug.Drug_id, &drug.Name, &drug.Unit_price); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		drugs = append(drugs, drug)
	}

	c.JSON(http.StatusOK, gin.H{"data": drugs})
}

// GetDrugById get drug from id specified in parameter
// and return as JSON
func GetDrugById(c *gin.Context) {
	db := models.SetupDB()

	var drug models.Drug
	sqlStatement := `SELECT * FROM drugs WHERE drug_id = $1`
	err := db.QueryRow(sqlStatement, c.Param("id")).Scan(&drug.Drug_id, &drug.Name, &drug.Unit_price)
	switch err {
	case sql.ErrNoRows:
		c.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
	case nil:
		c.JSON(http.StatusOK, gin.H{"data": drug})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

// EditDrug edit drug with id specified in parameter
// and return the status as JSON
func EditDrug(c *gin.Context) {
	db := models.SetupDB()

	var drug models.Drug
	sqlStatement := `SELECT * FROM drugs WHERE drug_id = $1`
	err := db.QueryRow(sqlStatement, c.Param("id")).Scan(&drug.Drug_id, &drug.Name, &drug.Unit_price)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	}

	var input models.Drug
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Name != nil {
		drug.Name = input.Name
	}
	if input.Unit_price != nil {
		drug.Unit_price = input.Unit_price
	}

	sqlStatement = `UPDATE drugs SET name=$1, unit_price=$2 WHERE drug_id=$3`
	_, err = db.Exec(sqlStatement, drug.Name, drug.Unit_price, c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "update success"})
}

// DeleteDrug deletes drug with id specified in parameter
// and return status as JSON
func DeleteDrug(c *gin.Context) {
	db := models.SetupDB()
	sqlStatement := `DELETE FROM drugs WHERE drug_id = $1`
	_, err := db.Exec(sqlStatement, c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "delete success"})
}
