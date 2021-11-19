package routes

import (
	"twistcode-test/controllers"

	"github.com/gorilla/mux"
)

func TransaksiRoutes(route *mux.Router) {
	route.HandleFunc("", controllers.GetAllTransaksi).Methods("GET")
	route.HandleFunc("/report", controllers.GetAllReportTransaksi).Methods("GET")
	route.HandleFunc("/getone/{id}", controllers.GetOneTransaksi).Methods("GET")
	route.HandleFunc("", controllers.CreateTransaksi).Methods("POST")
	route.HandleFunc("/edit/{id}", controllers.EditTransaksi).Methods("PUT")
	route.HandleFunc("/delete/{id}", controllers.DeleteTransaksi).Methods("DELETE")
}
