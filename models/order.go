package models

import "gorm.io/gorm"

type Order struct {
  gorm.Model
  UserID     uint
  User       User    `gorm:"foreignKey:UserID"`
  Status     string
  Total      float64
  Items      []OrderItem `gorm:"foreignKey:OrderID"` // one-to-many relationship with OrderItem
}