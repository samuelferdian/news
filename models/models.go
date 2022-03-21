package models

type News struct {
	Id     int    `json:"id" gorm:"primaryKey"`
	Title  string `json:"title" gorm:"coloum:title"`
	Desc   string `json:"desc" gorm:"coloum:desc"`
	Status string `json:"status" gorm:"coloum:status"`
}

type Tags struct {
	Id   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"coloum:name"`
}

type NewsTags struct {
	NewsId int `json:"news_id" gorm:"column:news_id"`
	TagId  int `json:"tag_id" gorm:"column:tag_id"`
}

type ParamInputNews struct {
	Title  string `json:"title" gorm:"coloum:title"`
	Desc   string `json:"desc" gorm:"coloum:desc"`
	Status string `json:"status" gorm:"coloum:status"`
	Tags   string `json:"tags" gorm:"coloum:tags"` // separate by comma
}

type ResultSearchNews struct {
	Title  string `json:"title" gorm:"coloum:title"`
	Desc   string `json:"desc" gorm:"coloum:desc"`
	Status string `json:"status" gorm:"coloum:status"`
	Name   string `json:"topic" gorm:"coloum:name"`
}
