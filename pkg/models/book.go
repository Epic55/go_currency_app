package models

import "gorm.io/gorm"

type Book struct {
	Id     int    `json:"id" gorm:"primaryKey"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"desc"`
}

type Rate struct {
	A_date string `xml:"date"`
	Items  []Item `xml:"item"`
}

type Item struct {
	Fullname    string `xml:"fullname"`
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Quant       string `xml:"quant"`
	Index       string `xml:"index"`
	Change      string `xml:"change"`
}

type RateModel struct {
	gorm.Model
	A_date string
	Item   []R_CURRENCY
}

type R_CURRENCY struct {
	gorm.Model
	RateModelID uint
	A_date      string
	Fullname    string
	Title       string
	Description string
	Quant       string
	Index       string
	Change      string
}

type Db_param struct {
	User     string
	Password string
	Host     string
	DbName   string
	Port     string
}
