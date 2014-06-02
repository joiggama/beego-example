package models

import (
	"time"
)

type Product struct {
	Id        int
	Code      string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (p Product) TableName() string {
  return "products"
}
