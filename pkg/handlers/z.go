package handlers

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Rate struct {
	//XMLName     xml.Name `xml:"rates"`
	Date  string `xml:"date"`
	Items []Item `xml:"item"`
}

type Item struct {
	//XMLName  xml.Name `xml:"item"`
	Fullname    string `xml:"fullname"`
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Quant       string `xml:"quant"`
	Index       string `xml:"index"`
	Change      string `xml:"change"`
}

type RateModel struct {
	gorm.Model
	Date string
	Item []ItemModel
}

type ItemModel struct {
	gorm.Model
	RateModelID uint
	Fullname    string
	Title       string
	Description string
	Quant       string
	Index       string
	Change      string
}

func (h handler) Z(w http.ResponseWriter, r *http.Request) {
	dbURL := "postgres://postgres:1@localhost:5432/db1"
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&RateModel{}, &ItemModel{})

	response, err := http.Get("https://nationalbank.kz/rss/get_rates.cfm?fdate=29.08.2023")
	if err != nil {
		fmt.Print(err.Error())
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var rate Rate
	err = xml.Unmarshal([]byte(responseData), &rate)
	if err != nil {
		log.Fatal("Error - ", err)
	}

	// Create a new RateModel instance
	rateModel := RateModel{
		Date: rate.Date,
	}

	// Convert and save items
	for _, item := range rate.Items {
		rateModel.Item = append(rateModel.Item, ItemModel{
			Fullname:    item.Fullname,
			Title:       item.Title,
			Description: item.Description,
			Quant:       item.Quant,
			Index:       item.Index,
			Change:      item.Change,
			Date:        rate.Date,
		})
	}

	db.Create(&rateModel)

	fmt.Println("Data saved successfully")
}
