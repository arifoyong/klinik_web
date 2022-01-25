package controllers

import (
	"net/http"
	"strings"

	"github.com/arifoyong/klinik/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type AddPatientVal struct {
	ID        uint   `json:"id" validate:"required"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	IC        string `json:"ic" validate:"required"`
	Email     string `json:"email" validate:"required"`
	Phone     string `json:"phone" validate:"required"`
	Address   string `json:"address"`
}

type EditPatientVal struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	IC        string `json:"ic"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
}

// GetPatients get information about all patients
// and return as JSON
func GetPatients(c *gin.Context) {
	var patients []models.Patient
	if err := models.DB.Find(&patients).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": patients})
}

// GetPatientByName find a patient with name given as parameter
// and return as JSON
func GetPatientByName(c *gin.Context) {
	var patients []models.Patient

	searchArg := "%" + strings.ToLower(c.Param("name")) + "%"
	if err := models.DB.Where("LOWER(firstname) LIKE ?", searchArg).Or("LOWER(lastname) LIKE ?", searchArg).Find(&patients).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": patients})
}

// GetPatientById find a patient by ID given in parameter
// and return it as JSON
func GetPatientById(c *gin.Context) {
	var patient models.Patient

	if err := models.DB.Where("id = ?", c.Param("id")).First(&patient).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": patient})
}

// AddPatient create a patient based on JSON parameters
// provided. If successful, patient will be returned
// as JSON
func AddPatient(c *gin.Context) {
	var input AddPatientVal
	var patient models.Patient

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()
	if err := validate.Struct(input); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErrors.Error()})
		return
	}

	patient.ID = input.ID
	patient.Firstname = input.Firstname
	patient.Lastname = input.Lastname
	patient.Email = input.Email
	patient.IC = input.IC
	patient.Address = input.Address
	patient.Phone = input.Phone

	if err := models.DB.Create(&patient).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": patient})
}

// EditPatient get patient based on ID given in parameter
// perform update & return the result as JSON
func EditPatient(c *gin.Context) {
	var input EditPatientVal
	var patient models.Patient

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.DB.Where("id = ?", c.Param("id")).First(&patient).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.DB.Model(&patient).Updates(input).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "update success"})
}

// DeletePatient find patient by ID given in parameter
// and return the status as JSON
func DeletePatient(c *gin.Context) {
	var patient models.Patient
	if err := models.DB.Where("id = ?", c.Param("id")).First(&patient).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Delete(&patient)
	c.JSON(http.StatusOK, gin.H{"data": "delete success"})
}
