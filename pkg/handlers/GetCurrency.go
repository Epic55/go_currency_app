package handlers

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	//"os"
)

type Rates struct {
	XMLName xml.Name `xml:"rates"`
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
	response, err := http.Get("https://nationalbank.kz/rss/get_rates.cfm?fdate=21.08.2023")
	if err != nil {
		fmt.Print(err.Error())
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var var1 Rates
	xml.Unmarshal(responseData, &var1)

	w.Header().Add("Content-Type", "application/xml")
	w.WriteHeader(http.StatusOK)
	xml.NewEncoder(w).Encode(var1)
}
