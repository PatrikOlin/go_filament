package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github.com/PatrikOlin/go_filament/models"
)

type updateBrandInput struct {
	Name string
}

func GetAllBrands(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var brands []models.Brand
	db.Find(&brands)

	c.JSON(http.StatusOK, gin.H{"data": brands})
}

func CreateBrand(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var input updateBrandInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	brand := models.Brand{
		Name: input.Name,
	}

	db.Create(&brand)

	c.JSON(http.StatusOK, gin.H{"data": brand})
}

func FindBrandByName(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var brand models.Brand
	if err := db.Where("name = ?", c.Param("name")).First(&brand).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Brand not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": brand})
}

func UpdateBrand(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var brand models.Brand
	if err := db.Where("name = ?", c.Param("name")).First(&brand).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Brand not found"})
		return
	}

	var input updateBrandInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	brand.Name = input.Name
	db.Save(&brand)

	c.JSON(http.StatusOK, gin.H{"data": brand})
}

func DeleteBrand(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var brand models.Brand
	if err := db.Where("name = ?", c.Param("name")).First(&brand).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Brand not found"})
		return
	}	

	db.Delete(&brand)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
