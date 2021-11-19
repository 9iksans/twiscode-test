package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"twistcode-test/helper"
	"twistcode-test/services"

	"github.com/gorilla/mux"
)

func ConvertB2D(w http.ResponseWriter, r *http.Request) {

	query := mux.Vars(r)
	splited := strings.Split(query["binary"], "")
	result := services.BinaryToDecimal(splited)
	w.Header().Set("Content-Type", "application/json")
	//Response Writer
	json.NewEncoder(w).Encode(helper.BuildSuccessResponse(result))

}

func ConvertD2B(w http.ResponseWriter, r *http.Request) {
	query := mux.Vars(r)
	value, _ := strconv.Atoi(query["decimal"])
	result := services.DecimalToBinary(value)
	w.Header().Set("Content-Type", "application/json")
	//Response Writer
	json.NewEncoder(w).Encode(helper.BuildSuccessResponse(result))

}
