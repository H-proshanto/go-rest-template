package svc

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
}
