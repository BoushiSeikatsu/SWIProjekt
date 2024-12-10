package controllers

import (
	"swi-warehouse/initializers"
	"swi-warehouse/models"

	"github.com/gin-gonic/gin"
)

func CheckIfManufacturerExists(manufacturerName string) bool {
	var manufacturer models.Manufacturer
	r := initializers.DB.First(&manufacturer, "name = ?", manufacturerName)
	if r.Error != nil {
		if r.Error.Error() == "record not found" {
			return false
		}
	}

	return true
}

func AllManufacturers() []models.Manufacturer {
	var manufacturers []models.Manufacturer
	initializers.DB.Find(&manufacturers)

	return manufacturers
}

func GetManufacturerByID(id string) models.Manufacturer {
	var manufacturer models.Manufacturer
	initializers.DB.First(&manufacturer, id)

	return manufacturer
}

func AddManufacturer(c *gin.Context) {
	manufacturerName := c.PostForm("name")

	if manufacturerName == "" {
		c.JSON(400, gin.H{"error": "Manufacturer name cannot be empty"})
		return
	}

	// check if manufacturer exists
	if CheckIfManufacturerExists(manufacturerName) {
		c.JSON(400, gin.H{"error": "Manufacturer already exists"})
		return
	}

	manufacturer := models.Manufacturer{
		Name: manufacturerName,
	}

	initializers.DB.Create(&manufacturer)

	c.Redirect(302, "/admin/addManufacturer")
}

func RenameManufacturer(c *gin.Context) {
	id := c.PostForm("manufacturerSelect")
	newManufacturerName := c.PostForm("name")

	if id == "" || newManufacturerName == "" {
		c.JSON(400, gin.H{"error": "Manufacturer name and new manufacturer name cannot be empty"})
		return
	}

	var manufacturer models.Manufacturer
	r := initializers.DB.First(&manufacturer, id)
	if r.Error != nil {
		c.JSON(400, gin.H{"error": "Manufacturer not found"})
		return
	}

	manufacturer.Name = newManufacturerName

	initializers.DB.Save(&manufacturer)

	c.Redirect(302, "/admin/renameManufacturer")
}

func RemoveManufacturer(c *gin.Context) {
	id := c.PostForm("manufacturerSelect")

	if id == "" {
		c.JSON(400, gin.H{"error": "Manufacturer name cannot be empty"})
		return
	}

	var manufacturer models.Manufacturer
	r := initializers.DB.First(&manufacturer, id)
	if r.Error != nil {
		c.JSON(400, gin.H{"error": "Manufacturer not found"})
		return
	}

	initializers.DB.Delete(&manufacturer)

	c.Redirect(302, "/admin/removeManufacturer")
}