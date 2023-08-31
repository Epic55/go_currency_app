package handlers

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	//"os"
)

type Rates struct {
	XMLName xml.Name `xml:"rates"`
	Date    string   `xml:"date"`
	Rates   []s1     `xml:"item"`
}

type s1 struct {
	XMLName     xml.Name `xml:"item"`
	Fullname    string   `xml:"fullname"`
	Title       string   `xml:"title"`
	Description string   `xml:"description"`
	Quant       int      `xml:"quant"`
	Index       int      `xml:"index"`
	Change      float64  `xml:"change"`
}

func (h handler) GetCurrency(w http.ResponseWriter, r *http.Request) {
	//defer r.Body.Close()

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

	var var1 Rates
	xml.Unmarshal(responseData, &var1)

	if err != nil {
		log.Fatalln(err)
	}

	// var var2 models.R_CURRENCY
	// xml.Unmarshal(responseData, &var2)

	// if result := h.DB.Create(&var2); result.Error != nil {
	// 	fmt.Println(result.Error)
	// }

	w.Header().Add("Content-Type", "application/xml")
	w.WriteHeader(http.StatusOK)
	xml.NewEncoder(w).Encode(var1)
}
