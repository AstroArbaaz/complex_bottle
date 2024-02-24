package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID						uint				`json:"productId" gorm:"primaryKey"`
	Name					string			`json:"name"`
	Description		string			`json:"description"`
	Details				string			`json:"details"`
	Quantity			float64			`json:"quantity" default:"0"`
	Price					float64			`json:"price" default:"0"`
	ProductImages 	[]string		`json:"product_images" gorm:"type:json"`
}