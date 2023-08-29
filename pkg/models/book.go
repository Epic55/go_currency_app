package models

type Book struct {
	Id     int    `json:"id" gorm:"primaryKey"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"desc"`
}

type R_CURRENCY struct {
	Id          int     `json:"id" gorm:"primaryKey"`
	Fullname    string  `xml:"fullname"`
	Title       string  `xml:"title"`
	Description string  `xml:"description"`
	Quant       int     `xml:"quant"`
	Index       int     `xml:"index"`
	Change      float64 `xml:"change"`
	A_date      string  `xml:"date"`
}
