package controllers

import (
	"strconv"
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
	initializers.DB.Preload("Manufacturer").Find(&products)

	return products
}

func AllProductsByManufacturer(manufacturerID string) []models.Product {
	var products []models.Product
	initializers.DB.Where("manufacturer_id= ?", manufacturerID).Find(&products)

	return products
}

func GetProductByID(id string) models.Product {
	var product models.Product
	initializers.DB.First(&product, id)

	return product
}

func AddProduct(c *gin.Context) {
	description := c.PostForm("description")
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

	// parse dimensions
	// string -> float64
	dx, err := strconv.ParseFloat(dimX, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid X dimension"})
		return
	}
	dy, err := strconv.ParseFloat(dimY, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid Y dimension"})
		return
	}
	dz, err := strconv.ParseFloat(dimZ, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid Z dimension"})
		return
	}

	// Find manufacturer
	var manufacturerObj models.Manufacturer
	r := initializers.DB.First(&manufacturerObj, manufacturer)
	if r.Error != nil {
		c.JSON(400, gin.H{"error": "Invalid manufacturer"})
		return
	}

	product := models.Product{
		Description: description,
		Manufacturer: manufacturerObj,
		Batch: batch,
		XDimension: dx,
		YDimension: dy,
		ZDimension: dz,
		Unit: unit,
	}

	initializers.DB.Create(&product)

	c.Redirect(302, "/admin/addProduct")
}

func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	description := c.PostForm("description")
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

	// parse dimensions
	// string -> float64
	dx, err := strconv.ParseFloat(dimX, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid X dimension"})
		return
	}
	dy, err := strconv.ParseFloat(dimY, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid Y dimension"})
		return
	}
	dz, err := strconv.ParseFloat(dimZ, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid Z dimension"})
		return
	}

	// Find manufacturer
	var manufacturerObj models.Manufacturer
	r := initializers.DB.First(&manufacturerObj, manufacturer)
	if r.Error != nil {
		c.JSON(400, gin.H{"error": "Invalid manufacturer"})
		return
	}

	var product models.Product
	r = initializers.DB.First(&product, id)
	if r.Error != nil {
		c.JSON(400, gin.H{"error": "Invalid product"})
		return
	}

	product.Description = description
	product.Manufacturer = manufacturerObj
	product.Batch = batch
	product.XDimension = dx
	product.YDimension = dy
	product.ZDimension = dz
	product.Unit = unit

	initializers.DB.Save(&product)

	c.Redirect(302, "/admin/searchProducts")
}

func SupplyProduct(c *gin.Context) {
	productID := c.PostForm("productSelect")
	storageID := c.PostForm("storageSelect")
	amount := c.PostForm("amount")

	if productID == "" || storageID == "" || amount == "" {
		c.JSON(400, gin.H{"error": "All fields are required"})
		return
	}

	// parse quantity
	q, err := strconv.Atoi(amount)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid quantity"})
		return
	}

	// Find product
	var product models.Product
	r := initializers.DB.First(&product, productID)
	if r.Error != nil {
		c.JSON(400, gin.H{"error": "Invalid product"})
		return
	}

	// Find storage
	var storage models.Storage
	r = initializers.DB.First(&storage, storageID)
	if r.Error != nil {
		c.JSON(400, gin.H{"error": "Invalid storage"})
		return
	}

	usr, _ := c.Get("user")

	supply := models.Added{
		Product:  product,
		Storage:  storage,
		Amount:   q,
		User: usr.(models.User),
	}

	initializers.DB.Create(&supply)

	// Find the stores with the product and storage
	var store models.Stores
	r = initializers.DB.Where("product_id = ? AND storage_id = ?", productID, storageID).First(&store)

	// update the store
	if r.Error == nil {
		store.Amount += q
		initializers.DB.Save(&store)
	} else {
		store := models.Stores{
			Product:  product,
			Storage:  storage,
			Amount: q,
		}
		initializers.DB.Create(&store)
	}

	c.Redirect(302, "/admin/selectSupplyManufacturer")
}

func WithdrawProduct(c *gin.Context) {
	productID := c.PostForm("productSelect")
	storageID := c.PostForm("storageSelect")
	amount := c.PostForm("amount")

	if productID == "" || storageID == "" || amount == "" {
		c.JSON(400, gin.H{"error": "All fields are required"})
		return
	}

	// parse quantity
	q, err := strconv.Atoi(amount)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid quantity"})
		return
	}

	// Find product
	var product models.Product
	r := initializers.DB.First(&product, productID)
	if r.Error != nil {
		c.JSON(400, gin.H{"error": "Invalid product"})
		return
	}

	// Find storage
	var storage models.Storage
	r = initializers.DB.First(&storage, storageID)
	if r.Error != nil {
		c.JSON(400, gin.H{"error": "Invalid storage"})
		return
	}

	// Find the stores with the product and storage
	var store models.Stores
	r = initializers.DB.Where("product_id = ? AND storage_id = ?", productID, storageID).First(&store)

	// check if there is enough product in the storage
	if r.Error != nil || store.Amount < q {
		c.JSON(400, gin.H{"error": "Not enough product in the storage"})
		return
	}
	
	// update the store
	if r.Error == nil {
		store.Amount -= q
		initializers.DB.Save(&store)
	}

	usr, _ := c.Get("user")

	remove := models.Removed{
		Product:  product,
		Storage:  storage,
		Amount:   q,
		User: usr.(models.User),
	}

	initializers.DB.Create(&remove)

	c.Redirect(302, "/admin/selectWithdrawManufacturer")
}