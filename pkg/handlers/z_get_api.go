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

	var rate models.Rate
	err = xml.Unmarshal([]byte(responseData), &rate)
	if err != nil {
		log.Fatal("Error - ", err)
	}

	// Create a new RateModel instance
	rateModel := models.RateModel{
		Date: rate.Date,
	}

	// Convert and save items
	for _, item := range rate.Items {
		rateModel.Item = append(rateModel.Item, models.R_CURRENCY{
			Fullname:    item.Fullname,
			Title:       item.Title,
			Description: item.Description,
			Quant:       item.Quant,
			Index:       item.Index,
			Change:      item.Change,
			//Date:        rate.Date,
		})
	}

	h.DB.Create(&rateModel)
	fmt.Println("Data saved successfully")

	w.Header().Add("Content-Type", "application/xml")
	w.WriteHeader(http.StatusOK)
	xml.NewEncoder(w).Encode("Done")
}
