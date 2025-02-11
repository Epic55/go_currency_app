package handlers

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"

	"github.com/Epic55/go_project_task/pkg/models"
	"github.com/gorilla/mux"
	log2 "github.com/sirupsen/logrus"
)

// GetAll 			godoc
// @Summary			Get All Currencirs from API.
// @Description		Return list of currencies.
// @Param			date1 path string true "Set date for currency"
// @Produce			application/xml
// @Tags			currency1
// @Success			200 {obejct} response.Response{}
// @Router			/currencys/date1 [get]
func (handler1 handler) Get_currency_from_api(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	date1, _ := vars["date1"]

	response, err := http.Get("https://nationalbank.kz/rss/get_rates.cfm?fdate=" + date1)
	if err != nil {
		log2.Error(err.Error())
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log2.Error(err)
	}

	var rate1 models.Rate
	err = xml.Unmarshal([]byte(responseData), &rate1)
	if err != nil {
		log2.Error("Error - ", err)
	}

	// Create a new RateModel instance
	ratemodel1 := models.RateModel{
		A_date: rate1.A_date,
	}

	// Convert and save items
	for _, i := range rate1.Items {
		ratemodel1.Item = append(ratemodel1.Item, models.R_CURRENCY{
			Title:  i.Title,
			Code:   i.Code,
			Value:  i.Value,
			A_date: rate1.A_date,
		})
	}

	var result2 []byte

	if result := handler1.DB.Create(&ratemodel1); result.Error != nil {
		fmt.Println(result.Error)
		result1 := map[string]bool{"success": false}
		result2, _ = json.Marshal(result1)
		log2.Info(string(result2))
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(string(result2))
	} else {
		result1 := map[string]bool{"success": true}
		result2, _ := json.Marshal(result1)
		log2.Info(string(result2))
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(string(result2))
	}

	// go func() {
	// 	if result := h.DB.Create(&ratemodel1); result.Error != nil {
	// 		fmt.Println(result.Error)
	// 		result1 := map[string]bool{"success": false}
	// 		result2, _ = json.Marshal(result1)
	// 		fmt.Println(string(result2))
	// 	} else {
	// 		result1 := map[string]bool{"success": true}
	// 		result2, _ := json.Marshal(result1)
	// 		fmt.Println("111 - ", string(result2))
	// 	}
	// }()
	// time.Sleep(time.Second)
	// w.Header().Add("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(string(result2))  //THIS GIVES EMPTY LINE IN POSTMAN
	// fmt.Println("222 - ", string(result2))

}
