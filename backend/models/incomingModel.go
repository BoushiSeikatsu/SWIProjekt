package models

import "gorm.io/gorm"

type Icoming struct {
	gorm.Model

	ProductID   uint `json:"productId"`
	Product		Product `json:"-"`
	Amount	   	float64 `json:"amount"`
	Suplier		string `json:"suplier"`
	IncomingID  string `json:"incomingId"`
	Note		string `json:"note"`
}
