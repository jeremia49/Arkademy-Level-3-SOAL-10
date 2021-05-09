package main

import (
	"arkademy/controllers"
	"arkademy/database"
	"arkademy/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB = database.InitDb()

func main() {
	server := gin.Default()
	server.LoadHTMLGlob("templates/*")

	api := server.Group("/api/")
	produkRepo := controllers.ControllerProduk(db)
	api.POST("/produk/create", produkRepo.CreateProdukJSON)
	api.GET("/produk", produkRepo.GetAllProduk)
	api.GET("/produk/:id", produkRepo.GetProduk)
	// api.PUT("/produk/:id", produkRepo.UpdateUser)
	// api.DELETE("/produk/:id", produkRepo.DeleteUser)

	web := server.Group("/")
	web.GET("/", func(c *gin.Context) {
		var data []models.Produk
		models.GetAllProduk(db, &data)
		c.HTML(http.StatusOK, "index.gohtml", gin.H{
			"title": "Home",
			"data":  data,
		})
	})

	web.GET("/create", func(c *gin.Context) {
		c.HTML(http.StatusOK, "create.gohtml", gin.H{
			"title": "Create Produk",
		})
	})

	web.POST("/create", func(c *gin.Context) {
		if err := controllers.CreateProdukForm(db, c); err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
		}
		c.Redirect(http.StatusFound, "/")
	})

	server.Run("localhost:8080")

}
