package controllers

import (
	"swi-warehouse/initializers"
	"swi-warehouse/models"

	"github.com/gin-gonic/gin"
)

func CheckIfStorageExists(storageName string) bool {
	var storage models.Storage
	r := initializers.DB.First(&storage, "name = ?", storageName)
	if r.Error != nil {
		if r.Error.Error() == "record not found" {
			return false
		}
	}

	return true
}

func AllStorages() []models.Storage {
	var storages []models.Storage
	initializers.DB.Find(&storages)

	return storages
}

func AddStorage(c *gin.Context) {
	storageName := c.PostForm("name")

	if storageName == "" {
		c.JSON(400, gin.H{"error": "Storage name cannot be empty"})
		return
	}

	// check if storage exists
	if CheckIfStorageExists(storageName) {
		c.JSON(400, gin.H{"error": "Storage already exists"})
		return
	}

	storage := models.Storage{
		Name: storageName,
	}

	initializers.DB.Create(&storage)

	c.Redirect(302, "/admin/addStorage")
}

func RenameStorage(c *gin.Context) {
	storageName := c.PostForm("storageSelect")
	newStorageName := c.PostForm("name")

	if storageName == "" || newStorageName == "" {
		c.JSON(400, gin.H{"error": "Storage name and new storage name cannot be empty"})
		return
	}

	var storage models.Storage
	r := initializers.DB.First(&storage, "name = ?", storageName)
	if r.Error != nil {
		c.JSON(400, gin.H{"error": "Storage not found"})
		return
	}

	storage.Name = newStorageName

	initializers.DB.Save(&storage)

	c.Redirect(302, "/admin/renameStorage")
}

func RemoveStorage(c *gin.Context) {
	storageName := c.PostForm("storageSelect")

	if storageName == "" {
		c.JSON(400, gin.H{"error": "Storage name cannot be empty"})
		return
	}

	var storage models.Storage
	r := initializers.DB.First(&storage, "name = ?", storageName)
	if r.Error != nil {
		c.JSON(400, gin.H{"error": "Storage not found"})
		return
	}

	initializers.DB.Delete(&storage)

	c.Redirect(302, "/admin/removeStorage")
}
