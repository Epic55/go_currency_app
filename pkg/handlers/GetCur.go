package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Epic55/go_project_task/pkg/models"
)

func (h handler) GetCur(w http.ResponseWriter, r *http.Request) {
	var books []models.R_CURRENCY

	if result := h.DB.Find(&books); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
}
