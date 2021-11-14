package models

import "time"

type Agenda struct {
	ID            int `gorm:"autoIncrement"`
	Name          string
	StartDateTime time.Time `gorm:"column:startDateTime"`
	EndDateTime   time.Time `gorm:"column:endDateTime"`
	Classes       string
	CreatedAt     time.Time `gorm:"column:createdAt"`
	UpdatedAt     time.Time `gorm:"column:updatedAt"`
}
