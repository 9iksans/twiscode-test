package main

import (
	"log"

	"net/http"

	"github.com/gorilla/mux"
	"twistcode-test/controllers"
	"twistcode-test/routes"
)

func main() {
	app := mux.NewRouter()
	app.HandleFunc("/api/convb2d/{binary}", controllers.ConvertB2D).Methods("GET")
	app.HandleFunc("/api/convd2b/{decimal}", controllers.ConvertD2B).Methods("GET")

	TransaksiRouter := app.PathPrefix("/api/transaksi").Subrouter()
	routes.TransaksiRoutes(TransaksiRouter)

	DetailTransaksiRouter := app.PathPrefix("/api/detail-transaksi").Subrouter()
	routes.DetailTransaksiRoutes(DetailTransaksiRouter)

	log.Fatal(http.ListenAndServe(":2000", app))
}
