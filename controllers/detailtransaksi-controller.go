package controllers

import (
	"encoding/json"
	"net/http"
	"twistcode-test/models"

	"github.com/gorilla/mux"
)

func GetOneDetailTransaksi(w http.ResponseWriter, r *http.Request) {
	// construct Models
	var detailTransaksi models.DetailTransaksi

	//Params & Body
	params := mux.Vars(r)

	//Models to ...
	result := detailTransaksi.GetOneDetailTransaksi(params["id"])

	//Response Header
	w.Header().Set("Content-Type", "application/json")
	if result.Code == 404 {
		w.WriteHeader(http.StatusNotFound)
	}
	//Response Writer
	json.NewEncoder(w).Encode(result)
}

func GetAllDetailTransaksi(w http.ResponseWriter, r *http.Request) {
	// construct Models
	var detailTransaksi models.DetailTransaksis

	//Models to ...
	result := detailTransaksi.GetAllDetailTransaksis()

	//Response Header
	w.Header().Set("Content-Type", "application/json")
	if result.Code == 500 {
		w.WriteHeader(http.StatusInternalServerError)
	}
	//Response Writer
	json.NewEncoder(w).Encode(result)

}

func CreateDetailTransaksi(w http.ResponseWriter, r *http.Request) {

	// construct Models
	var detailTransaksi models.DetailTransaksi

	//Params & Body
	json.NewDecoder(r.Body).Decode(&detailTransaksi)

	// Models to ...
	result := detailTransaksi.CreateDetailTransaksi()

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

func EditDetailTransaksi(w http.ResponseWriter, r *http.Request) {
	// construct Models
	var detailTransaksi models.DetailTransaksi

	//Params & Body
	json.NewDecoder(r.Body).Decode(&detailTransaksi)
	params := mux.Vars(r)

	// Models to ...
	result := detailTransaksi.EditDetailTransaksi(params["id"])

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

func DeleteDetailTransaksi(w http.ResponseWriter, r *http.Request) {
	// construct Models
	var detailTransaksi models.DetailTransaksi

	//Params & Body
	params := mux.Vars(r)

	// Models to ...
	result := detailTransaksi.DeleteDetailTransaksi(params["id"])

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
