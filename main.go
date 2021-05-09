package main

import (
	"arkademy/controllers"
	"arkademy/database"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	server := gin.Default()

	api := server.Group("/api/")

	var db *gorm.DB = database.InitDb()

	userRepo := controllers.ControllerProduk(db)
	api.POST("/produk/create", userRepo.CreateProduk)
	api.GET("/produk", userRepo.GetAllProduk)
	api.GET("/produk/:id", userRepo.GetProduk)
	// api.PUT("/produk/:id", userRepo.UpdateUser)
	// api.DELETE("/produk/:id", userRepo.DeleteUser)

	server.Run("localhost:8080")

}
