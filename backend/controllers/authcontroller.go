package controllers

import (
	"time"

	"github.com/BoushiSeikatsu/SWIProjekt/initializers"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Login(c *gin.Context) {
	// Export pass from the post data
	pass := c.PostForm("pass")

	// Check the pass against the current admin pass
	if pass != initializers.Config.Adminpass {
		// Respond with error and redirect to login page
		c.Redirect(302, "/login")
	}

	// Create JWT token and set it as a cookie
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"admin": true,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(initializers.Config.Secret))
	if err != nil {
		c.Redirect(302, "/login")
		return
	}

	c.SetCookie("auth", tokenString, 3600, "/", "", false, true)

	c.Redirect(302, "/")
}

func CheckLoggedIn(c *gin.Context) {
	// Get the token from the cookie
	token, err := c.Cookie("auth")
	if err != nil {
		c.Redirect(302, "/login")
		return
	}

	// Parse the token
	_, err = jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(initializers.Config.Secret), nil
	})
	if err != nil {
		c.Redirect(302, "/login")
		return
	}

	c.Next()
}
