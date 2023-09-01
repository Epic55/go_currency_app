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
	d, _ := vars["d"]

	var book models.R_CURRENCY

	if result := h.DB.First(&book, d); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}
