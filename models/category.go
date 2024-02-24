package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	ID uint `json:"categoryId"`
	Name string `json:"name"`
	Products []Product `gorm:"many2many:product_categories"`
}