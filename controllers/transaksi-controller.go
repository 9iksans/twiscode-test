package controllers

import (
	"encoding/json"
	"net/http"
	"twistcode-test/models"

	"github.com/gorilla/mux"
)

func GetOneTransaksi(w http.ResponseWriter, r *http.Request) {
	// construct Models
	var transaksi models.Transaksi

	//Params & Body
	params := mux.Vars(r)

	//Models to ...
	result := transaksi.GetOneTransaksi(params["id"])

	//Response Header
	w.Header().Set("Content-Type", "application/json")
	if result.Code == 404 {
		w.WriteHeader(http.StatusNotFound)
	}
	//Response Writer
	json.NewEncoder(w).Encode(result)
}

func GetAllTransaksi(w http.ResponseWriter, r *http.Request) {
	// construct Models
	var transaksi models.Transaksis

	//Models to ...
	result := transaksi.GetAllTransaksis()

	//Response Header
	w.Header().Set("Content-Type", "application/json")
	if result.Code == 500 {
		w.WriteHeader(http.StatusInternalServerError)
	}
	//Response Writer
	json.NewEncoder(w).Encode(result)

}

func GetAllReportTransaksi(w http.ResponseWriter, r *http.Request) {
	// construct Models
	var transaksi models.Transaksis
	var detailTransaksis models.DetailTransaksis

	// params := mux.Vars(r)
	//Models to ...
	result := transaksi.GetAllReportTransaksis(detailTransaksis)

	//Response Header
	w.Header().Set("Content-Type", "application/json")
	if result.Code == 500 {
		w.WriteHeader(http.StatusInternalServerError)
	}
	//Response Writer
	json.NewEncoder(w).Encode(result)

}

func CreateTransaksi(w http.ResponseWriter, r *http.Request) {

	// construct Models
	var transaksi models.Transaksi

	//Params & Body
	json.NewDecoder(r.Body).Decode(&transaksi)

	// Models to ...
	result := transaksi.CreateTransaksi()

	//Response Header
	w.Header().Set("Content-Type", "application/json")
	if result.Code == 403 {
		w.WriteHeader(http.StatusForbidden)
	} else if result.Code == 500 {
		w.WriteHeader(http.StatusInternalServerError)
	}
	//Response Writer
	json.NewEncoder(w).Encode(result)

}

func EditTransaksi(w http.ResponseWriter, r *http.Request) {
	// construct Models
	var transaksi models.Transaksi

	//Params & Body
	json.NewDecoder(r.Body).Decode(&transaksi)
	params := mux.Vars(r)

	// Models to ...
	result := transaksi.EditTransaksi(params["id"])

	//Response Header
	w.Header().Set("Content-Type", "application/json")
	if result.Code == 403 {
		w.WriteHeader(http.StatusForbidden)
	} else if result.Code == 404 {
		w.WriteHeader(http.StatusNotFound)
	} else if result.Code == 500 {
		w.WriteHeader(http.StatusInternalServerError)
	}
	//Response Writer
	json.NewEncoder(w).Encode(result)

}

func DeleteTransaksi(w http.ResponseWriter, r *http.Request) {
	// construct Models
	var transaksi models.Transaksi

	//Params & Body
	params := mux.Vars(r)

	// Models to ...
	result := transaksi.DeleteTransaksi(params["id"])

	//Response Header
	w.Header().Set("Content-Type", "application/json")
	if result.Code == 404 {
		w.WriteHeader(http.StatusNotFound)
	} else if result.Code == 500 {
		w.WriteHeader(http.StatusInternalServerError)
	}
	//Response Writer
	json.NewEncoder(w).Encode(result)

}
