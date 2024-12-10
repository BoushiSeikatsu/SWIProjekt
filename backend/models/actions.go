package models

import (
	"time"

	"gorm.io/gorm"
)

type Added struct {
	gorm.Model
	
	ProductID uint
	Product Product
	StorageID uint
	Storage Storage
	Date *time.Time
	Amount int
	UserID uint
	User User
}
