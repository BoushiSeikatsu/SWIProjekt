package models

import "gorm.io/gorm"

type Storage struct {
	gorm.Model

	Name string
}

type Stores struct {
	gorm.Model

	Amount int
	ProductID int
	Product Product
	StorageID int
	Storage Storage
}