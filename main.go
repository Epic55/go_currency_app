package main

import (
	"net/http"

	"github.com/Epic55/go_project_task/pkg/db"
	"github.com/Epic55/go_project_task/pkg/handlers"
	"github.com/Epic55/go_project_task/pkg/metric"
	"github.com/gorilla/mux"
	log2 "github.com/sirupsen/logrus"

	_ "github.com/Epic55/go_project_task/docs"
	httpSwagger "github.com/swaggo/http-swagger/v2"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// @title 	Currency Service API
// @version	1.0
// @description A Currency service API in Go using Gin framework

// @host 	localhost:8080
// @BasePath

func main() {
	DB := db.Init()
	h := handlers.New(DB)
	router := mux.NewRouter()
	router.Use(metric.PrometheusMiddleware)

	router.HandleFunc("/currency/{date1}", h.Get_currency_from_db).Methods(http.MethodGet)
	router.HandleFunc("/currency/{date1}/{code}", h.Get_currency_from_db).Methods(http.MethodGet)
	router.HandleFunc("/currencys/save/{date1}", h.Get_currency_from_api).Methods(http.MethodGet)

	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"), //The url pointing to API definition
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)
	router.Handle("/metrics", promhttp.Handler())

	metric.RecordMetrics()

	log2.Info("API is running!")
	http.ListenAndServe("127.0.0.1:8080", router)
}
