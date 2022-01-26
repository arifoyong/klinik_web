// package controllers

// import (
// 	"fmt"
// 	"net/http"
// 	"time"

// 	"github.com/arifoyong/klinik/models"
// 	"github.com/gin-gonic/gin"
// )

// type CreateNotaInput struct {
// 	ID             uint      `json:"id" binding:"required"`
// 	Date           time.Time `json:"date" binding:"required"`
// 	Nota_id        uint      `json:"nota_id" binding:"required"`
// 	Drug_name      string    `json:"drug_name"`
// 	Drug_type      string    `json:"drug_type"`
// 	Box_qty        uint      `json:"box_qty"`
// 	Vol_box        uint      `json:"vol_box"`
// 	Vol_unit       uint      `json:"vol_unit"`
// 	Box_cost       float32   `json:"box_cost"`
// 	Disc_percent   float32   `json:"disc_percent"`
// 	Disc_val       float32   `json:"disc_value"`
// 	Spdisc_percent float32   `json:"spdisc_percent"`
// 	Spdisc_val     float32   `json:"spdisc_value"`
// 	Ttl_cost_disc  float32   `json:"ttl_cost_disc"`
// 	Tax            float32   `json:"tax"`
// 	Ttl_cost       float32   `json:"ttl_cost"`
// 	Unit_cost      float32   `json:"unit_cost"`
// 	Vendor         string    `json:"vendor"`
// }

// type UpdateNotaInput struct {
// 	ID             uint      `json:"id"`
// 	Date           time.Time `json:"date"`
// 	Nota_id        uint      `json:"nota_id"`
// 	Drug_name      string    `json:"drug_name"`
// 	Drug_type      string    `json:"drug_type"`
// 	Box_qty        uint      `json:"box_qty"`
// 	Vol_box        uint      `json:"vol_box"`
// 	Vol_unit       uint      `json:"vol_unit"`
// 	Box_cost       float32   `json:"box_cost"`
// 	Disc_percent   float32   `json:"disc_percent"`
// 	Disc_val       float32   `json:"disc_value"`
// 	Spdisc_percent float32   `json:"spdisc_percent"`
// 	Spdisc_val     float32   `json:"spdisc_value"`
// 	Ttl_cost_disc  float32   `json:"ttl_cost_disc"`
// 	Tax            float32   `json:"tax"`
// 	Ttl_cost       float32   `json:"ttl_cost"`
// 	Unit_cost      float32   `json:"unit_cost"`
// 	Vendor         string    `json:"vendor"`
// }

// // ListNotas get all notas
// // and return it as JSON
// func ListNotas(c *gin.Context) {
// 	var notas []models.Nota
// 	models.DB.Find(&notas)

// 	c.JSON(http.StatusOK, gin.H{"data": notas})
// }

// // CreateNota creates a new nota based on
// // json input
// func CreateNota(c *gin.Context) {
// 	var inputs []CreateNotaInput
// 	var notas []models.Nota

// 	if err := c.ShouldBindJSON(&inputs); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	fmt.Println("Len of input", len(inputs))
// 	// Create nota
// 	for i, input := range inputs {
// 		fmt.Println(i)
// 		nota := models.Nota{ID: input.ID,
// 			Date:           input.Date,
// 			Nota_id:        input.Nota_id,
// 			Drug_name:      input.Drug_name,
// 			Drug_type:      input.Drug_type,
// 			Box_qty:        input.Box_qty,
// 			Vol_box:        input.Vol_box,
// 			Vol_unit:       input.Vol_unit,
// 			Box_cost:       input.Box_cost,
// 			Disc_percent:   input.Disc_percent,
// 			Disc_val:       input.Disc_val,
// 			Spdisc_percent: input.Spdisc_percent,
// 			Spdisc_val:     input.Spdisc_val,
// 			Ttl_cost_disc:  input.Ttl_cost_disc,
// 			Tax:            input.Tax,
// 			Ttl_cost:       input.Ttl_cost,
// 			Unit_cost:      input.Unit_cost,
// 			Vendor:         input.Vendor,
// 		}
// 		notas = append(notas, nota)
// 	}
// 	if err := models.DB.Create(&notas).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": notas})
// }

// // GetNotaById locate a nota whose ID matches
// // the ID parameter sent by client, then return
// // the result as JSON
// func GetNotaById(c *gin.Context) {
// 	var nota models.Nota

// 	if err := models.DB.Where("id = ?", c.Param("id")).First(&nota).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": nota})
// }

// // UpdateNota updates a nota whose ID matches
// // the ID parameter sent by client, then return
// // the result as JSON
// func UpdateNota(c *gin.Context) {
// 	var nota models.Nota
// 	if err := models.DB.Where("id = ?", c.Param("id")).First(&nota).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
// 		return
// 	}

// 	var input UpdateNotaInput
// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	if err := models.DB.Model(&nota).Updates(input).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"data": nota})
// }

// // DeleteNota finds a nota whose ID matches
// // the ID parameter sent by client then deletes it
// func DeleteNota(c *gin.Context) {
// 	var nota models.Nota
// 	if err := models.DB.Where("id = ?", c.Param("id")).First(&nota).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
// 		return
// 	}

// 	models.DB.Delete(&nota)
// 	c.JSON(http.StatusOK, gin.H{"data": true})
// }
