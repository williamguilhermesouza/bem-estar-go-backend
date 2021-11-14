package models

import "time"

type Attendance struct {
	ID             int       `gorm:"autoIncrement"`
	PatientId      int       `gorm:"column:patientId"`
	AttendanceDate time.Time `gorm:"column:attendanceDate"`
	DoneProcedures string    `gorm:"column:doneProcedures"`
	CreatedAt      time.Time `gorm:"column:createdAt"`
	UpdatedAt      time.Time `gorm:"column:updatedAt"`
}
