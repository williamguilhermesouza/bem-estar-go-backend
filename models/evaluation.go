package models

import "time"

type Evaluation struct {
	ID                     int    `gorm:"autoIncrement"`
	PatientId              int    `gorm:"column:patientId"`
	RightFeet              string `gorm:"column:rightFeet"`
	LeftFeet               string `gorm:"column:leftFeet"`
	RightAnkle             string `gorm:"column:rightAnkle"`
	LeftAnkle              string `gorm:"column:leftAnkle"`
	RightKnee              string `gorm:"column:rightKnee"`
	LeftKnee               string `gorm:"column:leftKnee"`
	Pelvis                 string
	Lumbar                 string
	Dorsal                 string
	Cervical               string
	RightShoulder          string `gorm:"column:rightShoulder"`
	LeftShoulder           string `gorm:"column:leftShoulder"`
	ShoulderBlade          string `gorm:"column:shoulderBlade"`
	Head                   string
	Observations           string
	PhysiotherapyDiagnosis string    `gorm:"column:physiotherapyDiagnosis"`
	RehabTarget            string    `gorm:"column:rehabTarget"`
	CreatedAt              time.Time `gorm:"column:createdAt"`
	UpdatedAt              time.Time `gorm:"column:updatedAt"`
}
