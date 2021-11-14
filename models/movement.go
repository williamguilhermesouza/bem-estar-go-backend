package models

import "time"

type Movement struct {
	ID          int `gorm:"autoIncrement"`
	Description string
	Value       float32
	CreatedAt   time.Time `gorm:"column:createdAt"`
	UpdatedAt   time.Time `gorm:"column:updatedAt"`
}
