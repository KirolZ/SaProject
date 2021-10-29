package entity

import (
	"time"

	"gorm.io/gorm"
)

type Nurse struct {
	gorm.Model

	Name  string
	Email string `gorm:"uniqueIndex"`
	Pass  string

	Screenings []Screening `gorm:"foreignKey:NurseID"`
}

type Disease struct {
	gorm.Model

	Name        string
	Description string

	Screenings []Screening `gorm:"foreignKey:DiseaseID"`
}

type MedicalRecord struct {
	gorm.Model

	Hospital_number string
	Personal_ID     string
	Patient_Name    string
	Patient_Age     int
	Patient_dob     time.Time
	Patient_Tel     string
	Screenings      []Screening `gorm:"foreignKey:MedRecID"`
}

type Screening struct {
	gorm.Model

	SaveTime time.Time

	Symptoms        string
	Weight          float32
	Height          float32
	Temperature     float32
	PulseRate       int
	RespirationRate int

	MedRecID *uint
	MedRec   MedicalRecord

	DiseaseID *uint
	Disease   Disease

	NurseID *uint
	Nurse   Nurse
}
