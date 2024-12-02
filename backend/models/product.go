package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Description string
	XDimension   float64
	YDimension   float64
	ZDimension   float64
	Batch		string
	Unit		string
	ManufacturerID int
	Manufacturer Manufacturer
}
