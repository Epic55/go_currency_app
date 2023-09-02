package models

import "encoding/xml"

type Book struct {
	Id     int    `json:"id" gorm:"primaryKey"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"desc"`
}

type Rates struct {
	XMLName xml.Name     `xml:"rates"`
	A_date  string       `xml:"date"`
	Rates   []R_CURRENCY `xml:"item"`
}

type R_CURRENCY struct {
	Id          int    `xml:"id" gorm:"primaryKey"`
	Fullname    string `xml:"fullname"`
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Quant       int    `xml:"quant"`
	Index       int    `xml:"index"`
	Change      string `xml:"change"`
	A_date      string `xml:"date"`
}
