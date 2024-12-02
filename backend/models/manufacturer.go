package models

import "gorm.io/gorm"

type Manufacturer struct {
	gorm.Model
	Name     string
	Products []Product
}
