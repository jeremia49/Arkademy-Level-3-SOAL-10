package models

import "gorm.io/gorm"

type Produk struct {
	*gorm.Model
	NamaProduk string
	Keterangan string
	Harga      string
	Jumlah     string
}

func CreateProduk(db *gorm.DB, Produk *Produk) (err error) {
	err = db.Create(&Produk).Error
	if err != nil {
		return err
	}
	return nil
}

func GetProduk(db *gorm.DB, Produk *Produk, id string) (err error) {
	err = db.Where("id = ?", id).Find(&Produk).Error
	if err != nil {
		return err
	}
	return nil
}

func GetAllProduk(db *gorm.DB, Produk *[]Produk) (err error) {
	err = db.Find(&Produk).Error
	if err != nil {
		return err
	}
	return nil
}
