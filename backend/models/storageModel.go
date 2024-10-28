package models

import "gorm.io/gorm"

type Storage struct {
	gorm.Model
	Name string `json:"name"`
}
