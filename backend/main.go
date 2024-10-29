package main

import (
	"github.com/BoushiSeikatsu/SWIProjekt/controllers"
	"github.com/BoushiSeikatsu/SWIProjekt/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.ConnectToDB()
	initializers.Load()
}

func main() {
	// Create a defaul router
	r := gin.Default()

	// TODO - Add cors handling
	
	// Routes
	r.GET("/login", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Login",
		})
	})

	r.GET("/", controllers.CheckLoggedIn, func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	r.GET("/index", controllers.CheckLoggedIn, func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	r.GET("/storage", controllers.CheckLoggedIn, func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Storage",
		})
	})

	r.Run(initializers.Config.IP + ":" + initializers.Config.Port)
}
