package models

import (
	"strconv"
	"twistcode-test/config"
	"twistcode-test/helper"
	// "github.com/go-playground/validator/v10"
	// "gorm.io/gorm"
)

type ReportTransaksi struct {
	ID                uint   `json:"id"`
	TanggalOrder      string `json:"tanggal_order"`
	StatusPelunasan   string `json:"status_pelunasan"`
	TanggalPembayaran string `json:"tanggal_pembayaran" `
	Jumlah            uint64 `json:"jumlah"`
	Subtotal          uint64 `json:"subtotal"`
}

type ReportTransaksis []ReportTransaksi

type Transaksi struct {
	ID                uint   `gorm:"primaryKey" json:"id"`
	TanggalOrder      string `json:"tanggal_order"`
	StatusPelunasan   string `json:"status_pelunasan"`
	TanggalPembayaran string `json:"tanggal_pembayaran" `
}

type Transaksis []Transaksi

func (transaksi Transaksi) GetOneTransaksi(id string) helper.Response {
	//autoMigrate and DB Connection
	db := config.ConnectToDatabase()
	db.AutoMigrate(&Transaksi{})

	//Queries
	data := db.Where("id = ?", id).First(&transaksi)

	//Return Response
	if data.Error != nil {
		return helper.BuildErrorResponse(404, "Not Found", nil)
	} else {
		return helper.BuildSuccessResponse(transaksi)
	}

}

func (transaksis Transaksis) GetAllTransaksis() helper.Response {

	//autoMigrate and DB Connection
	db := config.ConnectToDatabase()
	db.AutoMigrate(&Transaksi{})
	//Queries
	data := db.Find(&transaksis)

	//Return Response
	if data.Error != nil {
		return helper.BuildErrorResponse(500, "Internal Server Error!", nil)
	} else {
		return helper.BuildSuccessResponse(transaksis)
	}

}

func (transaksis Transaksis) GetAllReportTransaksis(detailTransaksis DetailTransaksis) helper.Response {

	//autoMigrate and DB Connection
	db := config.ConnectToDatabase()
	var reportTransaksis ReportTransaksis
	db.AutoMigrate(&Transaksi{})
	//Queries

	data := db.Model(&transaksis).Select("transaksis.id, transaksis.tanggal_order, transaksis.status_pelunasan, transaksis.tanggal_pembayaran, detail_transaksis.subtotal, detail_transaksis.jumlah").Joins("left join detail_transaksis on transaksis.id = detail_transaksis.id_transaksi").Find(&reportTransaksis)

	//Return Response
	if data.Error != nil {
		return helper.BuildErrorResponse(500, "Internal Server Error!", nil)
	} else {
		return helper.BuildSuccessResponse(reportTransaksis)
	}

}

func (transaksi Transaksi) CreateTransaksi() helper.Response {

	//autoMigrate and DB Connection
	db := config.ConnectToDatabase()
	db.AutoMigrate(&Transaksi{})

	//POST
	data := db.Create(&transaksi)

	//Return Response
	if data.Error != nil {
		return helper.BuildErrorResponse(500, "Internal Server Error", nil)
	} else {
		return helper.BuildSuccessResponse(transaksi)
	}

}

func (transaksi Transaksi) EditTransaksi(id string) helper.Response {
	//autoMigrate and DB Connection
	db := config.ConnectToDatabase()
	db.AutoMigrate(&Transaksi{})

	//PUT
	data := db.Model(&transaksi).Where("id = ?", id).Updates(transaksi)
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

func (transaksi Transaksi) DeleteTransaksi(id string) helper.Response {
	//autoMigrate and DB Connection
	db := config.ConnectToDatabase()
	db.AutoMigrate(&Transaksi{})

	//PUT
	data := db.Delete(&transaksi, id)
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
