package models

import (
//"time"
)

type Product struct {
	Id   int    `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
	//CreatedAt time.Time
	//UpdatedAt time.Time
}

func (p Product) TableName() string {
	return "products"
}
