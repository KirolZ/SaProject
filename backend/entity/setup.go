package entity

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("sa-64-project.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	database.AutoMigrate(
		&Nurse{}, &Disease{}, &MedicalRecord{}, &Screening{},
	)

	db = database

	MedicalRecord1 := MedicalRecord{
		Hospital_number: "1234",
		Personal_ID:     "21",
		Patient_Name:    "Oliver",
		Patient_Age:     23,
		Patient_dob:     time.Now(),
		Patient_Tel:     "09124214",
	}
	db.Model(&MedicalRecord{}).Create(&MedicalRecord1)

	Disease1 := Disease{
		Name:        "sore throat",
		Description: "Most sore throats are caused by infections",
	}
	db.Model(&Disease{}).Create(&Disease1)

	Nurse1 := Nurse{
		NurseName:  "pakapon",
		NurseEmail: "nueng@gmail.com",
		NursePass:  "12345",
	}
	db.Model(&Nurse{}).Create(&Nurse1)

}
