package controllers

import (
	"swi-warehouse/initializers"
	"swi-warehouse/models"

	"github.com/gin-gonic/gin"
)

func CheckIfProductExists(productName string) bool {
	var product models.Product
	r := initializers.DB.First(&product, "name = ?", productName)
	if r.Error != nil {
		if r.Error.Error() == "record not found" {
			return false
		}
	}

	return true
}

func AllProducts() []models.Product {
	var products []models.Product
	initializers.DB.Find(&products)

	return products
}

func AddProduct(c *gin.Context) {
	/*description := c.PostForm("description")
	manufacturer := c.PostForm("manufacturerSelect")
	batch := c.PostForm("batch")
	dimX := c.PostForm("dimX")
	dimY := c.PostForm("dimY")
	dimZ := c.PostForm("dimZ")
	unit := c.PostForm("unit")

	if description == "" || manufacturer == "" || batch == "" || dimX == "" || dimY == "" || dimZ == "" || unit == "" {
		c.JSON(400, gin.H{"error": "All fields are required"})
		return
	}

	// Find manufacturer


	product := models.Product{
		Description: description,
		Manufacturer: manufacturer,
		Batch: batch,
		XDimension: dimX,
		DimY: dimY,
		DimZ: dimZ,
		Unit: unit,
	}*/
}