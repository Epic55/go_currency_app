package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Epic55/go_project_task/pkg/models"
	"github.com/gorilla/mux"
	log2 "github.com/sirupsen/logrus"
)

func (h handler) Get_currency_from_db(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	A_date, _ := vars["date1"]
	Code, _ := vars["code"]

	var cur []models.R_CURRENCY
	if Code == "" {
		if result := h.DB.Where("a_date = ?", A_date).Find(&cur); result.Error != nil {
			log2.Error(result.Error)
		} else {
			log2.Info("Search is done")
		}
	} else {
		if result := h.DB.Where("a_date = ? AND code = ?", A_date, Code).Find(&cur); result.Error != nil {
			log2.Error(result.Error)
		} else {
			log2.Info("Search is done")
		}
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cur)
}
