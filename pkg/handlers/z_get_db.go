package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Epic55/go_project_task/pkg/models"
	"github.com/gorilla/mux"
)

func (h handler) Db(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	A_date, _ := vars["date1"]
	var cur models.R_CURRENCY

	if result := h.DB.Where("A_date <> ?", A_date).Find(&cur); result.Error != nil {
		fmt.Println(result.Error)
	} else {
		fmt.Println("Search is done")
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cur)
}
