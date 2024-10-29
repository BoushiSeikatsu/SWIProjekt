package models

import "gorm.io/gorm"

type Distribution struct {
	gorm.Model

	ProductID uint    `json:"productId"`
	Product   Product `json:"-"`
	Amount    float64 `json:"amount"`
	Reciever  string  `json:"reciever"`
	Note      string  `json:"note"`
	RecieptID string  `json:"recieptId"`
}
