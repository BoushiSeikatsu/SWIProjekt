package models

import "gorm.io/gorm"

type Store struct {
	gorm.Model
	
	ProductID   uint    `json:"tileId"`
	Product     Product `json:"-"`
	StorageID 	uint    `json:"storageId"`
	Storage   	Storage `json:"-"`
	Unit      	string  `json:"unit"`
	Quantity  	float64 `json:"quantity"`
}
