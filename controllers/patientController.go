package controllers

import (
	"database/sql"
	"net/http"
	"strings"
	"time"

	"github.com/arifoyong/klinik/models"
	"github.com/gin-gonic/gin"
)

type AddPatientVal struct {
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	IC        string    `json:"ic" validate:"required"`
	DOB       time.Time `json:"dob" validate:"required"`
	Email     string    `json:"email" validate:"required"`
	Phone     string    `json:"phone" validate:"required"`
	Address   string    `json:"address"`
}

type EditPatientVal struct {
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	IC        string    `json:"ic"`
	DOB       time.Time `json:"dob"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Address   string    `json:"address"`
}

// GetPatients get information about all patients
// and return as JSON
func GetPatients(c *gin.Context) {
	db := models.SetupDB()

	rows, err := db.Query("SELECT * FROM patients")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var patients []models.Patient
	for rows.Next() {
		var id uint
		var firstname string
		var lastname string
		var ic string
		var dob time.Time
		var email string
		var phone string
		var address string

		err = rows.Scan(&id, &firstname, &lastname, &ic, &dob, &email, &phone, &address)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		newPatient := models.Patient{
			ID:        id,
			Firstname: firstname,
			Lastname:  lastname,
			IC:        ic,
			DOB:       dob,
			Email:     email,
			Phone:     phone,
			Address:   address,
		}

		patients = append(patients, newPatient)
	}

	c.JSON(http.StatusOK, gin.H{"data": patients})

}

// GetPatientByName find a patient with name given as parameter
// and return as JSON
func GetPatientByName(c *gin.Context) {
	db := models.SetupDB()

	arg := "%" + strings.ToLower(c.Param("name")) + "%"
	sqlStatement := `SELECT * FROM patients WHERE LOWER(firstname) LIKE $1 OR LOWER(lastname) LIKE $1`
	rows, err := db.Query(sqlStatement, arg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var patients []models.Patient
	for rows.Next() {
		var id uint
		var firstname string
		var lastname string
		var ic string
		var dob time.Time
		var email string
		var phone string
		var address string

		err = rows.Scan(&id, &firstname, &lastname, &ic, &dob, &email, &phone, &address)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		newPatient := models.Patient{
			ID:        id,
			Firstname: firstname,
			Lastname:  lastname,
			IC:        ic,
			DOB:       dob,
			Email:     email,
			Phone:     phone,
			Address:   address,
		}
		patients = append(patients, newPatient)
	}

	c.JSON(http.StatusOK, gin.H{"data": patients})
}

// GetPatientById find a patient by ID given in parameter
// and return it as JSON
func GetPatientById(c *gin.Context) {
	var patient models.Patient
	db := models.SetupDB()

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

// // AddPatient create a patient based on JSON parameters
// // provided. If successful, patient will be returned
// // as JSON
// func AddPatient(c *gin.Context) {
// 	var input AddPatientVal
// 	var patient models.Patient

// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	validate := validator.New()
// 	if err := validate.Struct(input); err != nil {
// 		validationErrors := err.(validator.ValidationErrors)
// 		c.JSON(http.StatusBadRequest, gin.H{"error": validationErrors.Error()})
// 		return
// 	}

// 	patient.Firstname = input.Firstname
// 	patient.Lastname = input.Lastname
// 	patient.Email = input.Email
// 	patient.IC = input.IC
// 	patient.DOB = input.DOB
// 	patient.Address = input.Address
// 	patient.Phone = input.Phone

// 	if err := models.DB.Create(&patient).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": patient})
// }

// // EditPatient get patient based on ID given in parameter
// // perform update & return the result as JSON
// func EditPatient(c *gin.Context) {
// 	var input EditPatientVal
// 	var patient models.Patient

// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	if err := models.DB.Where("id = ?", c.Param("id")).First(&patient).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	if err := models.DB.Model(&patient).Updates(input).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": "update success"})
// }

// // DeletePatient find patient by ID given in parameter
// // and return the status as JSON
// func DeletePatient(c *gin.Context) {
// 	var patient models.Patient
// 	if err := models.DB.Where("id = ?", c.Param("id")).First(&patient).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	models.DB.Delete(&patient)
// 	c.JSON(http.StatusOK, gin.H{"data": "delete success"})
// }
