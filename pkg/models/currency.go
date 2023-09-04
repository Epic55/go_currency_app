package models

import "gorm.io/gorm"

type Rate struct {
	A_date string `xml:"date"`
	Items  []Item `xml:"item"`
}

type Item struct {
	Title string `xml:"fullname"`
	Code  string `xml:"title"`
	Value string `xml:"description"`
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
	Title       string
	Code        string
	Value       string
}

type Db_param struct {
	User     string
	Password string
	Host     string
	DbName   string
	Port     string
}
