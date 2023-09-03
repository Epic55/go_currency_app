package handlers

import (
	"encoding/json"
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
	date1, _ := vars["date1"]

	response, err := http.Get("https://nationalbank.kz/rss/get_rates.cfm?fdate=" + date1)
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

	var result2 []byte
	if result := h.DB.Create(&v1); result.Error != nil {
		fmt.Println(result.Error)
		result1 := map[string]bool{"success": false}
		result2, _ = json.Marshal(result1)
		fmt.Println(string(result2))
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(string(result2))
	} else {
		fmt.Println("Data saved successfully")
		result1 := map[string]bool{"success": true}
		result2, _ := json.Marshal(result1)
		fmt.Println(string(result2))
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(string(result2))
	}
	//go result, err := h.DB.Create(&v1)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//w.Header().Add("Content-Type", "application/xml")
	//w.WriteHeader(http.StatusOK)
	//xml.NewEncoder(w).Encode("Done")
	//time.Sleep(time.Second)
}
