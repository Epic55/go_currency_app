package main

import (
	"log"
	"net/http"

	"github.com/Epic55/go_project_task/pkg/db"
	"github.com/Epic55/go_project_task/pkg/handlers"
	"github.com/gorilla/mux"
)

func main() {
	DB := db.Init()
	h := handlers.New(DB)
	router := mux.NewRouter()

	router.HandleFunc("/currency/{date1}", h.Get_currency_from_db).Methods(http.MethodGet)
	router.HandleFunc("/currency/{date1}/{code}", h.Get_currency_from_db).Methods(http.MethodGet)
	router.HandleFunc("/currencys/{date1}", h.Get_currency_from_api).Methods(http.MethodGet)

	log.Println("API is running!")
	http.ListenAndServe("127.0.0.1:4000", router)
}
