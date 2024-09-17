package main

import (
	"awesomeProject/controller"
	"awesomeProject/dao"
	"awesomeProject/dbproperties"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {
	db := dbproperties.ConnectToDB()
	dao.SetDB(db)
	r := mux.NewRouter()
	r.HandleFunc("/product", controller.Handler).Methods(http.MethodPost, http.MethodPut)
	r.HandleFunc("/product/{id:[0-9]+}", controller.Handler).Methods(http.MethodGet, http.MethodDelete)
	r.HandleFunc("/product/{name:[a-z]+}", controller.Handler).Methods(http.MethodGet, http.MethodDelete)
	r.HandleFunc("/products", controller.Handler).Methods(http.MethodGet)

	err := http.ListenAndServe(":8000", r)
	if err != nil {
		log.Fatal(err)
	}
}
