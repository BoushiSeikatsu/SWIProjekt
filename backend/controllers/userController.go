package controllers

import (
	"swi-warehouse/initializers"
	"swi-warehouse/models"

	"github.com/gin-gonic/gin"
)

func CheckIfUserExists (username string) bool {
	var user models.User
	r := initializers.DB.First(&user, "username = ?", username)
	if r.Error != nil {
		if r.Error.Error() == "record not found" {
			return false
		}
	}

	return true
}

func AddUser (c *gin.Context) {
	userName := c.PostForm("username")

	if userName == "" {
		c.JSON(400, gin.H{"error": "Username cannot be empty"})
		return
	}

	// check if user exists
	if CheckIfUserExists(userName) {
		c.JSON(400, gin.H{"error": "User already exists"})
		return
	}

	user := models.User{
		Username: userName,
	}

	initializers.DB.Create(&user)

	c.Redirect(302, "/admin/addUser")
}

func ChangePassword (c *gin.Context) {
	id := c.PostForm("userSelect")
	password := c.PostForm("newPassword")

	if id == "" || password == "" {
		c.JSON(400, gin.H{"error": "Username and password cannot be empty"})
		return
	}

	var user models.User
	r := initializers.DB.First(&user, id)
	if r.Error != nil {
		c.JSON(400, gin.H{"error": r.Error.Error()})
		return
	}

	user.Password = password
	initializers.DB.Save(&user)

	c.Redirect(302, "/admin/changePassword")
}

func DisableUser (c *gin.Context) {
	id := c.PostForm("userSelect")

	if id == "" {
		c.JSON(400, gin.H{"error": "User cannot be empty"})
		return
	}

	var user models.User
	r := initializers.DB.First(&user, id)
	if r.Error != nil {
		c.JSON(400, gin.H{"error": "User not found"})
		return
	}

	user.Active = false
	initializers.DB.Save(&user)

	c.Redirect(302, "/admin/disableUser")
}

func EnableUser (c *gin.Context) {
	id := c.PostForm("userSelect")

	if id == "" {
		c.JSON(400, gin.H{"error": "User cannot be empty"})
		return
	}

	var user models.User
	r := initializers.DB.First(&user, id)
	if r.Error != nil {
		c.JSON(400, gin.H{"error": "User not found"})
		return
	}

	user.Active = true
	initializers.DB.Save(&user)

	c.Redirect(302, "/admin/disableUser")
}

func AllUsers () []models.User {
	var users []models.User
	initializers.DB.Find(&users)

	return users
}

func AllActiveUsers() []models.User {
	var users []models.User
	initializers.DB.Find(&users, "active = ?", true)

	return users
}

func AllDisabledUsers() []models.User {
	var users []models.User
	initializers.DB.Find(&users, "active = ?", false)

	return users
}
