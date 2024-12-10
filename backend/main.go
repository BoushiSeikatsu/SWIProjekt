package main

import (
	"net/http"
	"swi-warehouse/controllers"
	"swi-warehouse/gintemplrenderer"
	"swi-warehouse/initializers"
	"swi-warehouse/templates"
	"swi-warehouse/templates/admin/user"
	"swi-warehouse/templates/admin/manufacturer"
	"swi-warehouse/templates/admin/storage"
	"swi-warehouse/templates/admin/product"

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
            c.HTML(http.StatusOK, "", templatesAdminUser.AdminAddUser())
        })

        // User thingies
        admin.GET("/addUser", controllers.AdminMiddleware(), func (c *gin.Context) {
            c.HTML(http.StatusOK, "", templatesAdminUser.AdminAddUser())
        })
        admin.POST("/addUser", controllers.AdminMiddleware(), controllers.AddUser)

        admin.GET("/changePassword", controllers.AdminMiddleware(), func (c *gin.Context) {
            c.HTML(http.StatusOK, "", templatesAdminUser.AdminChangePassword(controllers.AllUsers()))
        })
        admin.POST("/changePassword", controllers.AdminMiddleware(), controllers.ChangePassword)

        admin.GET("/disableUser", controllers.AdminMiddleware(), func (c *gin.Context) {
            c.HTML(http.StatusOK, "", templatesAdminUser.AdminDisableUser(controllers.AllActiveUsers()))
        })
        admin.POST("/disableUser", controllers.AdminMiddleware(), controllers.DisableUser)

        admin.GET("/enableUser", controllers.AdminMiddleware(), func (c *gin.Context) {
            c.HTML(http.StatusOK, "", templatesAdminUser.AdminEnableUser(controllers.AllDisabledUsers()))
        })
        admin.POST("/enableUser", controllers.AdminMiddleware(), controllers.EnableUser)
        
        // Storage thingies
        admin.GET("/addStorage", controllers.AdminMiddleware(), func (c *gin.Context) {
            c.HTML(http.StatusOK, "", templatesAdminStorage.AdminAddStorage())
        })
        admin.POST("/addStorage", controllers.AdminMiddleware(), controllers.AddStorage)

        admin.GET("/renameStorage", controllers.AdminMiddleware(), func (c *gin.Context) {
            c.HTML(http.StatusOK, "", templatesAdminStorage.AdminRenameStorage(controllers.AllStorages()))
        })
        admin.POST("/renameStorage", controllers.AdminMiddleware(), controllers.RenameStorage)

        admin.GET("/removeStorage", controllers.AdminMiddleware(), func (c *gin.Context) {
            c.HTML(http.StatusOK, "", templatesAdminStorage.AdminRemoveStorage(controllers.AllStorages()))
        })
        admin.POST("/removeStorage", controllers.AdminMiddleware(), controllers.RemoveStorage)

        // Manufacturer thingies
        admin.GET("/addManufacturer", controllers.AdminMiddleware(), func (c *gin.Context) {
            c.HTML(http.StatusOK, "", templatesAdminManufacturer.AdminAddManufacturer())
        })
        admin.POST("/addManufacturer", controllers.AdminMiddleware(), controllers.AddManufacturer)

        admin.GET("/renameManufacturer", controllers.AdminMiddleware(), func (c *gin.Context) {
            c.HTML(http.StatusOK, "", templatesAdminManufacturer.AdminRenameManufacturer(controllers.AllManufacturers()))
        })
        admin.POST("/renameManufacturer", controllers.AdminMiddleware(), controllers.RenameManufacturer)

        admin.GET("/removeManufacturer", controllers.AdminMiddleware(), func (c *gin.Context) {
            c.HTML(http.StatusOK, "", templatesAdminManufacturer.AdminRemoveManufacturer(controllers.AllManufacturers()))
        })
        admin.POST("/removeManufacturer", controllers.AdminMiddleware(), controllers.RemoveManufacturer)

        // Product thingies
        admin.GET("/addProduct", controllers.AdminMiddleware(), func (c *gin.Context) {
            c.HTML(http.StatusOK, "", templatesAdminProduct.AdminAddProduct(controllers.AllManufacturers()))
        })
        admin.POST("/addProduct", controllers.AdminMiddleware(), controllers.AddProduct)

        admin.GET("/locateProduct/:id", controllers.AdminMiddleware(), func (c *gin.Context) {
            c.HTML(http.StatusOK, "", templatesAdminProduct.AdminLocateProduct(controllers.GetStores(c.Param("id"))))
        })

        admin.GET("/updateProduct/:id", controllers.AdminMiddleware(), func (c *gin.Context) {
            c.HTML(http.StatusOK, "", templatesAdminProduct.AdminUpdateProduct(controllers.GetProductByID(c.Param("id")), controllers.AllManufacturers()))
        })
        admin.POST("/updateProduct/:id", controllers.AdminMiddleware(), controllers.UpdateProduct)

        admin.GET("/searchProducts", controllers.AdminMiddleware(), func (c *gin.Context) {
            c.HTML(http.StatusOK, "", templatesAdminProduct.AdminSearchProduct(controllers.AllProducts()))
        })

        // Product add remove thingies
        admin.GET("/selectSupplyManufacturer", controllers.AdminMiddleware(), func (c *gin.Context) {
            c.HTML(http.StatusOK, "", templatesAdminProduct.AdminSelectSupplyManufacturer(controllers.AllManufacturers()))
        })
        admin.POST("/selectSupplyManufacturer", controllers.AdminMiddleware(), func (c *gin.Context) {
            // parse manufacturer id
            c.Redirect(http.StatusFound, "/admin/selectSupplyProduct/" + c.PostForm("manufacturerSelect"))
        })

        admin.GET("/selectSupplyProduct/:id", controllers.AdminMiddleware(), func (c *gin.Context) {
            c.HTML(http.StatusOK, "", templatesAdminProduct.AdminSelectSupplyProduct(controllers.GetManufacturerByID(c.Param("id")), controllers.AllProductsByManufacturer(c.Param("id")), controllers.AllStorages()))
        })
        admin.POST("/selectSupplyProduct/:id", controllers.AdminMiddleware(), controllers.SupplyProduct)
    
        // Product add remove thingies
        admin.GET("/selectWithdrawManufacturer", controllers.AdminMiddleware(), func (c *gin.Context) {
            c.HTML(http.StatusOK, "", templatesAdminProduct.AdminSelectWithdrawManufacturer(controllers.AllManufacturers()))
        })
        admin.POST("/selectWithdrawManufacturer", controllers.AdminMiddleware(), func (c *gin.Context) {
            // parse manufacturer id
            c.Redirect(http.StatusFound, "/admin/selectWithdrawProduct/" + c.PostForm("manufacturerSelect"))
        })
    
        admin.GET("/selectWithdrawProduct/:id", controllers.AdminMiddleware(), func (c *gin.Context) {
            c.HTML(http.StatusOK, "", templatesAdminProduct.AdminSelectWithdrawProduct(controllers.GetManufacturerByID(c.Param("id")), controllers.AllProductsByManufacturer(c.Param("id")), controllers.AllStorages()))
        })
        admin.POST("/selectWithdrawProduct/:id", controllers.AdminMiddleware(), controllers.WithdrawProduct)
    
        admin.GET("/productHistory", controllers.AdminMiddleware(), controllers.ProductHistory)
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
