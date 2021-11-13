package models

import "time"

type User struct {
	ID             int `gorm:"autoIncrement"`
	Name           string
	Sex            string
	PhoneNumber    string `gorm:"column:phoneNumber"`
	StreetName     string `gorm:"column:streetName"`
	StreetNumber   string `gorm:"column:streetNumber"`
	StreetDistrict string `gorm:"column:streetDistrict"`
	City           string
	State          string
	BirthDate      time.Time `gorm:"column:birthDate"`
	Cpf            string
	Email          string
	Password       string
	CreatedAt      time.Time `gorm:"column:createdAt"`
	UpdatedAt      time.Time `gorm:"column:updatedAt"`
}
