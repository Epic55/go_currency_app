package handlers

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Epic55/go_project_task/pkg/models"
	"github.com/gorilla/mux"
)

func (h handler) GetCurrency(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	vars := mux.Vars(r)
	d, _ := vars["d"]

	response, err := http.Get("https://nationalbank.kz/rss/get_rates.cfm?fdate=" + d)
	if err != nil {
		fmt.Print(err.Error())
	}
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var var1 models.Rates
	xml.Unmarshal(responseData, &var1)

	if err != nil {
		log.Fatalln(err)
	}

	//SAVE TO DB
	var var2 models.R_CURRENCY
	xml.Unmarshal(responseData, &var2)

	if result := h.DB.Create(&var2); result.Error != nil {
		fmt.Println(result.Error)
	}
	fmt.Println(responseData)
	w.Header().Add("Content-Type", "application/xml")
	w.WriteHeader(http.StatusOK)
	xml.NewEncoder(w).Encode(var1)
}
