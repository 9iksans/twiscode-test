package models

import (
	"strconv"
	"twistcode-test/config"
	"twistcode-test/helper"
	// "github.com/go-playground/validator/v10"
	// "gorm.io/gorm"
)

type DetailTransaksi struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	IDTransaksi uint   `json:"id_transaksi"`
	Harga       uint64 `json:"harga"`
	Jumlah      uint64 `json:"jumlah"`
	Subtotal    uint64 `json:"subtotal"`
}

type DetailTransaksis []DetailTransaksi

func (detailTransaksi DetailTransaksi) GetOneDetailTransaksi(id string) helper.Response {
	//autoMigrate and DB Connection
	db := config.ConnectToDatabase()
	db.AutoMigrate(&DetailTransaksi{})

	//Queries
	data := db.Where("id = ?", id).First(&detailTransaksi)

	//Return Response
	if data.Error != nil {
		return helper.BuildErrorResponse(404, "Not Found", nil)
	} else {
		return helper.BuildSuccessResponse(detailTransaksi)
	}

}

func (detailTransaksis DetailTransaksis) GetAllDetailTransaksis() helper.Response {

	//autoMigrate and DB Connection
	db := config.ConnectToDatabase()
	db.AutoMigrate(&DetailTransaksi{})
	//Queries
	data := db.Find(&detailTransaksis)

	//Return Response
	if data.Error != nil {
		return helper.BuildErrorResponse(500, "Internal Server Error!", nil)
	} else {
		return helper.BuildSuccessResponse(detailTransaksis)
	}

}

func (detailTransaksi DetailTransaksi) CreateDetailTransaksi() helper.Response {

	//autoMigrate and DB Connection
	db := config.ConnectToDatabase()
	db.AutoMigrate(&DetailTransaksi{})

	//POST
	detailTransaksi.Subtotal = detailTransaksi.Harga * detailTransaksi.Jumlah
	data := db.Create(&detailTransaksi)

	//Return Response
	if data.Error != nil {
		return helper.BuildErrorResponse(500, "Internal Server Error", nil)
	} else {
		return helper.BuildSuccessResponse(detailTransaksi)
	}

}

func (detailTransaksi DetailTransaksi) EditDetailTransaksi(id string) helper.Response {
	//autoMigrate and DB Connection
	db := config.ConnectToDatabase()
	db.AutoMigrate(&DetailTransaksi{})

	//PUT
	detailTransaksi.Subtotal = detailTransaksi.Harga * detailTransaksi.Jumlah
	data := db.Model(&detailTransaksi).Where("id = ?", id).Updates(detailTransaksi)
	conv, _ := strconv.Atoi(id)

	//Return Response
	if data.Error != nil {
		return helper.BuildErrorResponse(500, "Internet Server Error", nil)
	} else if data.RowsAffected == 0 {
		return helper.BuildErrorResponse(404, "Not Found", nil)
	} else {
		return helper.BuildSuccessResponse(map[string]interface{}{"id": conv, "update": "Successful"})
	}
}

func (detailTransaksi DetailTransaksi) DeleteDetailTransaksi(id string) helper.Response {
	//autoMigrate and DB Connection
	db := config.ConnectToDatabase()
	db.AutoMigrate(&DetailTransaksi{})

	//PUT
	data := db.Delete(&detailTransaksi, id)
	conv, _ := strconv.Atoi(id)

	//Return Response
	if data.Error != nil {
		return helper.BuildErrorResponse(500, "Internal Server Error", nil)
	} else if data.RowsAffected == 0 {
		return helper.BuildErrorResponse(404, "Not Found", nil)
	} else {
		return helper.BuildSuccessResponse(map[string]interface{}{"id": conv, "delete": "successful"})
	}
}
