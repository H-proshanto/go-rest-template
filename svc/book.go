package svc

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Name string `json:"name"`
	ISBN int64  `json:"isbn"`
}
