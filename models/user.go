package models

import (
	"time"

	"gorm.io/gorm"
)


type User struct {
  gorm.Model
  ID        	uint            `gorm:"primaryKey" json:"id"`
  Username  	*string         `gorm:"unique" json:"username"`
	Fullname		string					`json:"full_name"`
  Email     	*string         `gorm:"unique"`
  Password  	string					`json:"password"`
  Address   	*string					`json:"address"`
  Phone    		*string					`json:"phone"`
	Birthday 		*time.Time			`json:"DOB"`
	CreatedAt    time.Time      `json:"createdAt"`// Automatically managed by GORM for creation time
  UpdatedAt    time.Time      `json:"updatedAt"`// Automatically managed by GORM for update time÷
	Activated 	bool						`json:"isActivated" gorm:"default:false"`
	Deleted			bool						`json:"isDeleted" gorm:"default:false"`
}