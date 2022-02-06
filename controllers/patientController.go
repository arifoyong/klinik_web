package controllers

import (
	"database/sql"
	"net/http"
	"strings"

	"github.com/arifoyong/klinik/models"
	"github.com/gin-gonic/gin"
)

// GetPatients get information about all patients
// and return as JSON
func GetPatients(c *gin.Context) {
	db := models.SetupDB()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM patients")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var patients []models.Patient
	for rows.Next() {
		var patient models.Patient
		err = rows.Scan(&patient.ID, &patient.Firstname, &patient.Lastname, &patient.IC, &patient.DOB, &patient.Email, &patient.Phone, &patient.Address)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		patients = append(patients, patient)
	}

	c.JSON(http.StatusOK, gin.H{"data": patients})

}

// GetPatientByName find a patient with name given as parameter
// and return as JSON
func GetPatientByName(c *gin.Context) {
	db := models.SetupDB()
	defer db.Close()

	arg := "%" + strings.ToLower(c.Param("name")) + "%"
	sqlStatement := `SELECT * FROM patients WHERE LOWER(firstname) LIKE $1 OR LOWER(lastname) LIKE $1`
	rows, err := db.Query(sqlStatement, arg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var patients []models.Patient
	for rows.Next() {
		var patient models.Patient
		err = rows.Scan(&patient.ID, &patient.Firstname, &patient.Lastname, &patient.IC, &patient.DOB, &patient.Email, &patient.Phone, &patient.Address)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		patients = append(patients, patient)
	}

	c.JSON(http.StatusOK, gin.H{"data": patients})
}

// GetPatientById find a patient by ID given in parameter
// and return it as JSON
func GetPatientById(c *gin.Context) {
	var patient models.Patient
	db := models.SetupDB()
	defer db.Close()

	sqlStatement := `SELECT * FROM patients WHERE id=$1`
	switch err := db.QueryRow(sqlStatement, c.Param("id")).Scan(
		&patient.ID,
		&patient.Firstname,
		&patient.Lastname,
		&patient.IC,
		&patient.DOB,
		&patient.Email,
		&patient.Phone,
		&patient.Address); err {
	case sql.ErrNoRows:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
	case nil:
		c.JSON(http.StatusOK, gin.H{"data": patient})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

// AddPatient create a patient based on JSON parameters
// provided. If successful, patient will be returned
// as JSON
func AddPatient(c *gin.Context) {
	var input models.Patient
	db := models.SetupDB()
	defer db.Close()

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// validate := validator.New()
	// if err := validate.Struct(input); err != nil {
	// 	validationErrors := err.(validator.ValidationErrors)
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": validationErrors.Error()})
	// 	return
	// }

	sqlStatement := `INSERT INTO patients (firstname, lastname, ic, dob, email, phone, address)
						VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`

	_, err := db.Exec(sqlStatement, input.Firstname, input.Lastname, input.IC, input.DOB, input.Email, input.Phone, input.Address)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "success"})
}

// EditPatient get patient based on ID given in parameter
// perform update & return the result as JSON
func EditPatient(c *gin.Context) {
	var patient models.Patient
	db := models.SetupDB()
	defer db.Close()

	sqlStatement := `SELECT * FROM patients WHERE id=$1`
	if err := db.QueryRow(sqlStatement, c.Param("id")).Scan(
		&patient.ID,
		&patient.Firstname,
		&patient.Lastname,
		&patient.IC,
		&patient.DOB,
		&patient.Email,
		&patient.Phone,
		&patient.Address); err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	}

	var input models.Patient
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	if input.Firstname != nil {
		patient.Firstname = input.Firstname
	}
	if input.Lastname != nil {
		patient.Lastname = input.Lastname
	}

	if input.IC != nil {
		patient.IC = input.IC
	}
	if input.DOB != nil {
		patient.DOB = input.DOB
	}
	if input.Email != nil {
		patient.Email = input.Email
	}
	if input.Phone != nil {
		patient.Phone = input.Phone
	}
	if input.Address != nil {
		patient.Address = input.Address
	}

	sqlStatement = `UPDATE patients SET firstname=$1, lastname=$2, ic=$3, dob=$4, email=$5, address=$6, phone=$7 WHERE id=$8`
	_, err := db.Exec(sqlStatement, patient.Firstname, patient.Lastname, patient.IC, patient.DOB, patient.Email, patient.Address, patient.Phone, c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "update success"})
}

// DeletePatient find patient by ID given in parameter
// and return the status as JSON
func DeletePatient(c *gin.Context) {
	db := models.SetupDB()
	defer db.Close()

	sqlStatement := `DELETE FROM patients where id = $1`
	_, err := db.Exec(sqlStatement, c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "delete success"})
}
