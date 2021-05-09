package main

import (
	"arkademy/controllers"
	"arkademy/database"
	"arkademy/models"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB = database.InitDb()

type IndexPage struct {
	Num    uint
	Produk models.Produk
}

func main() {
	server := gin.Default()
	server.LoadHTMLGlob("templates/*")

	api := server.Group("/api/")
	produkRepo := controllers.ControllerProduk(db)
	api.POST("/produk/create", produkRepo.CreateProdukJSON)
	api.GET("/produk", produkRepo.GetAllProdukJSON)
	api.GET("/produk/:id", produkRepo.GetProdukJSON)
	// api.PUT("/produk/:id", produkRepo.UpdateUser)
	// api.DELETE("/produk/:id", produkRepo.DeleteUser)

	web := server.Group("/")
	web.GET("/", func(c *gin.Context) {
		var data []models.Produk

		models.GetAllProduk(db, &data)

		var ix []IndexPage

		for i, el := range data {
			ix = append(ix, IndexPage{
				Num:    uint(i + 1),
				Produk: el,
			})
		}

		c.HTML(http.StatusOK, "index.gohtml", gin.H{
			"title": "Home",
			"data":  ix,
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

	web.GET("/delete/:id", func(c *gin.Context) {
		c.HTML(http.StatusOK, "delete.gohtml", gin.H{
			"title": "Hapus Produk",
		})
	})

	web.POST("/delete/:id", func(c *gin.Context) {
		if err := controllers.DeleteProdukForm(db, c); err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
		}
		c.Redirect(http.StatusFound, "/")
	})

	web.GET("/edit/:id", func(c *gin.Context) {

		produk, err := controllers.GetProduk(db, c)

		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.AbortWithStatus(http.StatusNotFound)
			}
			c.AbortWithStatus(http.StatusInternalServerError)
		}

		c.HTML(http.StatusOK, "edit.gohtml", gin.H{
			"title": "Edit Produk",
			"data":  produk,
		})
	})

	web.POST("/edit/:id", func(c *gin.Context) {

		if err := controllers.UpdateProduk(db, c); err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.AbortWithStatus(http.StatusNotFound)
			}
			c.AbortWithStatus(http.StatusInternalServerError)
		}

		c.Redirect(http.StatusFound, "/")
	})

	server.Run("localhost:8080")

}
