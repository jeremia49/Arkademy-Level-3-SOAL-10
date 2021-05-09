package controllers

import (
	"arkademy/models"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type produkRepo struct {
	Db *gorm.DB
}

func ControllerProduk(db *gorm.DB) *produkRepo {
	db.AutoMigrate(&models.Produk{})
	return &produkRepo{Db: db}
}

func (repository *produkRepo) CreateProdukJSON(c *gin.Context) {
	var produk models.Produk
	c.BindJSON(&produk)
	err := models.CreateProduk(repository.Db, &produk)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
}

func (repository *produkRepo) GetProduk(c *gin.Context) {
	id, _ := c.Params.Get("id")
	var produk models.Produk
	err := models.GetProduk(repository.Db, &produk, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, produk)
}

func (repository *produkRepo) GetAllProduk(c *gin.Context) {
	var produk []models.Produk
	err := models.GetAllProduk(repository.Db, &produk)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, produk)
}

func CreateProdukForm(Db *gorm.DB, c *gin.Context) (r error) {

	if c.PostForm("namaproduk") == "" || c.PostForm("keterangan") == "" || c.PostForm("harga") == "" || c.PostForm("jumlah") == "" {
		return gin.Error{}
	}

	produk := models.Produk{
		NamaProduk: c.PostForm("namaproduk"),
		Keterangan: c.PostForm("keterangan"),
		Harga:      c.PostForm("harga"),
		Jumlah:     c.PostForm("jumlah"),
	}

	err := models.CreateProduk(Db, &produk)
	if err != nil {
		return err
	}

	return nil
}
