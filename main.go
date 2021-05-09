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
	userRepo := controllers.ControllerProduk(db)
	api.POST("/produk/create", userRepo.CreateProduk)
	api.GET("/produk", userRepo.GetAllProduk)
	api.GET("/produk/:id", userRepo.GetProduk)
	// api.PUT("/produk/:id", userRepo.UpdateUser)
	// api.DELETE("/produk/:id", userRepo.DeleteUser)

	web := server.Group("/")
	web.GET("/", func(c *gin.Context) {

		var data []models.Produk

		models.GetAllProduk(db, &data)
		c.HTML(http.StatusOK, "index.gohtml", gin.H{
			"title": "Home",
			"data":  data,
		})
	})

	server.Run("localhost:8080")

}
