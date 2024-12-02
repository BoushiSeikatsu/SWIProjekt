package main

import (
	"net/http"
	"swi-warehouse/controllers"
	"swi-warehouse/gintemplrenderer"
	"swi-warehouse/initializers"
	"swi-warehouse/templates"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.Load()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

    r.SetTrustedProxies(nil)

    ginHtmlRenderer := r.HTMLRender
    r.HTMLRender = &gintemplrenderer.HTMLTemplRenderer{FallbackHtmlRenderer: ginHtmlRenderer}

	r.Use(CORSMiddleware())
    // First time setup
    r.GET("/setup", func (c *gin.Context) {
            if controllers.CheckIfAdminExists() {
                c.JSON(400, gin.H{"error": "Admin already exists"})
            } else {
                c.HTML(http.StatusOK, "", templates.SetupAdmin())
            }
        })
    r.POST("/setup", controllers.SetupAdmin)

    r.GET("/login", func (c *gin.Context) {
        c.HTML(http.StatusOK, "", templates.Login())
    })
    r.POST("/login", controllers.Login)

    r.GET("/", controllers.AuthMiddleware())

    admin := r.Group("/admin")
    {
        admin.GET("/", controllers.AdminMiddleware(), func (c *gin.Context) {
            c.HTML(http.StatusOK, "", templates.AdminAddUser())
        })

        // User thingies
        admin.GET("/addUser", controllers.AdminMiddleware(), func (c *gin.Context) {
            c.HTML(http.StatusOK, "", templates.AdminAddUser())
        })
        admin.POST("/addUser", controllers.AdminMiddleware(), controllers.AddUser)

        admin.GET("/changePassword", controllers.AdminMiddleware(), func (c *gin.Context) {
            c.HTML(http.StatusOK, "", templates.AdminChangePassword(controllers.AllUsers()))
        })
        admin.POST("/changePassword", controllers.AdminMiddleware(), controllers.ChangePassword)

        admin.GET("/disableUser", controllers.AdminMiddleware(), func (c *gin.Context) {
            c.HTML(http.StatusOK, "", templates.AdminDisableUser(controllers.AllActiveUsers()))
        })
        admin.POST("/disableUser", controllers.AdminMiddleware(), controllers.DisableUser)

        admin.GET("/enableUser", controllers.AdminMiddleware(), func (c *gin.Context) {
            c.HTML(http.StatusOK, "", templates.AdminEnableUser(controllers.AllDisabledUsers()))
        })
        admin.POST("/enableUser", controllers.AdminMiddleware(), controllers.EnableUser)
        
        // Storage thingies
        admin.GET("/addStorage", controllers.AdminMiddleware(), func (c *gin.Context) {
            c.HTML(http.StatusOK, "", templates.AdminAddStorage())
        })
        admin.POST("/addStorage", controllers.AdminMiddleware(), controllers.AddStorage)

        admin.GET("/renameStorage", controllers.AdminMiddleware(), func (c *gin.Context) {
            c.HTML(http.StatusOK, "", templates.AdminRenameStorage(controllers.AllStorages()))
        })
        admin.POST("/renameStorage", controllers.AdminMiddleware(), controllers.RenameStorage)

        admin.GET("/removeStorage", controllers.AdminMiddleware(), func (c *gin.Context) {
            c.HTML(http.StatusOK, "", templates.AdminRemoveStorage(controllers.AllStorages()))
        })
        admin.POST("/removeStorage", controllers.AdminMiddleware(), controllers.RemoveStorage)

        // Manufacturer thingies
        admin.GET("/addManufacturer", controllers.AdminMiddleware(), func (c *gin.Context) {
            c.HTML(http.StatusOK, "", templates.AdminAddManufacturer())
        })
        admin.POST("/addManufacturer", controllers.AdminMiddleware(), controllers.AddManufacturer)

        admin.GET("/renameManufacturer", controllers.AdminMiddleware(), func (c *gin.Context) {
            c.HTML(http.StatusOK, "", templates.AdminRenameManufacturer(controllers.AllManufacturers()))
        })
        admin.POST("/renameManufacturer", controllers.AdminMiddleware(), controllers.RenameManufacturer)

        admin.GET("/removeManufacturer", controllers.AdminMiddleware(), func (c *gin.Context) {
            c.HTML(http.StatusOK, "", templates.AdminRemoveManufacturer(controllers.AllManufacturers()))
        })
        admin.POST("/removeManufacturer", controllers.AdminMiddleware(), controllers.RemoveManufacturer)
    }

	r.Run(initializers.IP + ":" + initializers.PORT)
}

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
		orig := c.Request.Header.Clone().Get("Origin")
        c.Writer.Header().Set("Access-Control-Allow-Origin", orig)
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(http.StatusNoContent)
            return
        }

        c.Next()
    }
}
