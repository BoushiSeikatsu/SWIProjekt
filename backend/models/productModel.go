package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model

	Description    string  `json:"description"`
	ManufacturerID   string  `json:"manufacturerId"`
	ProductID      string  `json:"productId"`
	XSize          float64 `json:"xSize"`
	YSize          float64 `json:"ySize"`
	ZSize          float64 `json:"zSize"`
	Batch          string  `json:"batch"`
	Calibration    string  `json:"calibration"`
	Unit      	string  `json:"unit"`
	Image          string  `json:"image"`
}
