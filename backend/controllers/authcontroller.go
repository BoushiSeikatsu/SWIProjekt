package controllers

import (
	"swi-warehouse/initializers"
	"swi-warehouse/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func CheckIfAdminExists() bool {
	var user models.User
	r := initializers.DB.First(&user, "admin = ?", true)
	if r.Error != nil {
		if r.Error.Error() == "record not found" {
			return false
		}
	}
	return true
}

func SetupAdmin(c *gin.Context) {
	username := c.PostForm("username")

	if !CheckIfAdminExists() {
		user := models.User{
			Username:    username,
			Password: c.PostForm("password"),
			Admin:    true,
		}
		initializers.DB.Create(&user)
		c.Redirect(302, "/login")
		return
	}

	c.JSON(400, gin.H{"error": "Admin already exists"})
}

func Login (c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	var user models.User
	r := initializers.DB.First(&user, "username = ? AND password = ?", username, password)
	if r.Error != nil {
		c.JSON(400, gin.H{"error": "Invalid username or password"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"admin": user.Admin,
	})

	tokenString, err := token.SignedString([]byte(initializers.SECRET))

	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to generate token"})
		return
	}

	c.SetCookie("token", tokenString, 3600, "/", "", false, true)

	c.Redirect(302, "/")
}

func Logout (c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "", false, true)
	c.Redirect(302, "/login")
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("token")
		if err != nil {
			c.Redirect(302, "/login")
			return
		}

		t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			return []byte(initializers.SECRET), nil
		})

		if err != nil {
			c.Redirect(302, "/login")
			return
		}

		if t.Valid {
			c.Set("username", t.Claims.(jwt.MapClaims)["username"])
			c.Next()
			return
		}

		c.Redirect(302, "/login")
	}
}

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("token")
		if err != nil {
			c.Redirect(302, "/login")
			return
		}

		t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			return []byte(initializers.SECRET), nil
		})

		if err != nil {
			c.Redirect(302, "/login")
			return
		}

		if t.Valid {
			claims := t.Claims.(jwt.MapClaims)
			if claims["admin"] == true {
				c.Set("username", claims["username"])
				c.Set("admin", claims["admin"])
				c.Next()
				return
			}
		}

		c.Redirect(302, "/login")
	}
}