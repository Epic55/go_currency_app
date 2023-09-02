package handlers

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/Epic55/go_project_task/pkg/models"
	"github.com/gorilla/mux"
)

func (h handler) Api(w http.ResponseWriter, r *http.Request) {

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

	var v2 models.Rate
	err = xml.Unmarshal([]byte(responseData), &v2)
	if err != nil {
		log.Fatal("Error - ", err)
	}

	// Create a new RateModel instance
	v1 := models.RateModel{
		A_date: v2.A_date,
	}

	// Convert and save items
	for _, i := range v2.Items {
		v1.Item = append(v1.Item, models.R_CURRENCY{
			Fullname:    i.Fullname,
			Title:       i.Title,
			Description: i.Description,
			Quant:       i.Quant,
			Index:       i.Index,
			Change:      i.Change,
			A_date:      v2.A_date,
		})
	}

	if result := h.DB.Create(&v1); result.Error != nil {
		fmt.Println(result.Error)
	}
	fmt.Println("Data saved successfully")

	w.Header().Add("Content-Type", "application/xml")
	w.WriteHeader(http.StatusOK)
	go xml.NewEncoder(w).Encode("Done")
	time.Sleep(time.Second)
}
