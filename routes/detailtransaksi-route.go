package routes

import (
	"twistcode-test/controllers"

	"github.com/gorilla/mux"
)

func DetailTransaksiRoutes(route *mux.Router) {
	route.HandleFunc("", controllers.GetAllDetailTransaksi).Methods("GET")
	route.HandleFunc("/{id}", controllers.GetOneDetailTransaksi).Methods("GET")
	route.HandleFunc("", controllers.CreateDetailTransaksi).Methods("POST")
	route.HandleFunc("/{id}", controllers.EditDetailTransaksi).Methods("PUT")
	route.HandleFunc("/{id}", controllers.DeleteDetailTransaksi).Methods("DELETE")
}
