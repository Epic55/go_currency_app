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

	// for i := 0; i < len(var1.Rates); i++ {
	// 	fmt.Println("Fullname: " + var1.Rates[i].Fullname)
	// 	fmt.Println("Title: " + var1.Rates[i].Title)
	// 	fmt.Println("Description: " + var1.Rates[i].Description)
	// 	fmt.Println("Quant: ", var1.Rates[i].Quant)
	// 	fmt.Println("Index: ", var1.Rates[i].Index)
	// 	fmt.Println("Change: ", var1.Rates[i].Change, "\n")
	// }

	// if result := h.DB.Find(&var1); result.Error != nil {
	// 	fmt.Println(result.Error)
	// }

	w.Header().Add("Content-Type", "application/xml")
	w.WriteHeader(http.StatusOK)
	xml.NewEncoder(w).Encode(var1)
}
