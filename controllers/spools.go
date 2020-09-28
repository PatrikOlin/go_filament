package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github.com/PatrikOlin/go_filament/models"
	"github.com/PatrikOlin/go_filament/util"
)

// type CreateSpoolInput struct {
// 	Brand string `json:"brand" binding:"required"`
// 	Name string `json:"name" binding:"required"`
// 	Weight int `json:"weight"`
// 	SpoolWeight int `json:"spool_weight"`
// 	Color string `json:"color"`
// 	Material string `json:"material"`
// 	NozzleTemp int `json:"nozzle_temp"`
// 	PlateTemp int `json:"plate_temp"`
// 	PricePerKg int `json:"price_per_kg"`
// 	Superpowers []models.Superpower `json:"superpowers"`
// 	Notes string `json:"notes"`
// }

type UpdateSpoolInput struct {
	Brand       string   `json:"brand" `
	Name        string   `json:"name" `
	Weight      int      `json:"weight"`
	SpoolWeight int      `json:"spool_weight"`
	Color       string   `json:"color"`
	Material    string   `json:"material"`
	NozzleTemp  int      `json:"nozzle_temp"`
	PlateTemp   int      `json:"plate_temp"`
	PricePerKg  int      `json:"price_per_kg"`
	Superpowers []string `json:"superpowers"`
	Notes       string   `json:"notes"`
}

func GetAllSpools(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var spools []models.Spool
	db.Preload("Brand").Preload("Superpowers").Find(&spools)

	c.JSON(http.StatusOK, gin.H{"data": spools})
}

func CreateSpool(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var input UpdateSpoolInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var brand models.Brand
	if res := db.Where("name = ?", input.Brand).FirstOrCreate(&brand,
		models.Brand{Name: input.Brand}); res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": res.Error})
		return
	}

	var powers []models.Superpower
	for _, pwr := range input.Superpowers {
		var pow models.Superpower
		if err := db.Where("name = ?", pwr).FirstOrCreate(&pow,
			models.Superpower{Name: pwr}); err.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
			return
		}
		powers = append(powers, pow)
	}

	spool := models.Spool{
		Tag:         util.GenerateTag(),
		Brand:       brand,
		Name:        input.Name,
		Weight:      input.Weight,
		SpoolWeight: input.SpoolWeight,
		Color:       input.Color,
		Material:    input.Material,
		NozzleTemp:  input.NozzleTemp,
		PlateTemp:   input.PlateTemp,
		PricePerKg:  input.PricePerKg,
		Notes:       input.Notes,
		Superpowers: powers,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   nil,
	}

	if err := db.Create(&spool).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": spool})
}

func FindSpoolById(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var spool models.Spool
	if err := db.Where("id = ?", c.Param("id")).First(&spool).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Spool not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": spool})
}

func FindSpoolByTag(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var spool models.Spool
	if err := db.Preload("Brand").Preload("Superpowers").Where("tag = ?", c.Param("tag")).First(&spool).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Spool not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": spool})
}

func UpdateSpool(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var spool models.Spool
	if err := db.Where("tag = ?", c.Param("tag")).First(&spool).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Spool not found"})
		return
	}

	var input UpdateSpoolInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&spool).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": spool})
}

func DeleteSpool(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var spool models.Spool
	if err := db.Where("tag = ?", c.Param("tag")).First(&spool).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Spool not found"})
		return
	}

	db.Delete(&spool)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
