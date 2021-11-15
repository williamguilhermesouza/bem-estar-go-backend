package models

import "time"

type Evaluation struct {
	ID                         int    `gorm:"autoIncrement"`
	PatientId                  int    `gorm:"column:patientId"`
	CurrentDiseaseHistoric     string `gorm:"column:currentDiseaseHistoric"`
	AssociateDiseases          string `gorm:"column:associateDiseases"`
	PastDiseases               string `gorm:"column:pastDiseases"`
	FamilyHistoric             string `gorm:"column:familyHistoric"`
	LifeHabits                 string `gorm:"column:lifeHabits"`
	ClinicalDiagnosis          string `gorm:"column:clinicalDiagnosis"`
	MainComplain               string `gorm:"column:mainComplain"`
	Medication                 string
	Mobility                   string
	ConsciousState             string `gorm:"column:consciousState"`
	SkinAndMucous              string `gorm:"column:skinAndMucous"`
	AccessWays                 string `gorm:"column:accessWays"`
	ThoraxFormat               string `gorm:"column:thoraxFormat"`
	VentilationType            string `gorm:"column:ventilationType"`
	VentilationMuscularPattern string `gorm:"column:ventilationMuscularPattern"`
	VentilationRhythm          string `gorm:"column:ventilationRhythm"`
	Abdomen                    string
	Signals                    string
	Symptoms                   string
	ThoracicMobility           string  `gorm:"column:thoracicMobility"`
	LungExpansion              string  `gorm:"column:lungExpansion"`
	RespiratoryFrequency       float32 `gorm:"column:respiratoryFrequency"`
	CardiacFrequency           float32 `gorm:"column:cardiacFrequency"`
	Saturation                 float32
	ArterialPressure           float32 `gorm:"column:arterialPressure"`
	Temperature                float32
	LungHearing                string `gorm:"column:lungHearing"`
	Cough                      string
	Secretion                  string
	VentilationMode            string `gorm:"column:ventilationMode"`
	Tonus                      string
	ReflexesAndReaction        string `gorm:"column:reflexesAndReaction"`
	MuscularStrength           string `gorm:"column:muscularStrength"`
	Sensibility                string
	MotorControl               string `gorm:"column:motorControl"`
	ComplimentaryExams         string `gorm:"column:complimentaryExams"`
	FunctionalKineticDiagnosis string `gorm:"column:functionalKineticDiagnosis"`
	Problem                    string
	Target                     string
	Conduct                    string
	CreatedAt                  time.Time `gorm:"column:createdAt"`
	UpdatedAt                  time.Time `gorm:"column:updatedAt"`
}
