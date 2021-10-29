package entity

import (
	"time"

	"golang.org/x/crypto/bcrypt"

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

	MedicalRecord2 := MedicalRecord{
		Hospital_number: "1214",
		Personal_ID:     "22",
		Patient_Name:    "Hanza",
		Patient_Age:     25,
		Patient_dob:     time.Now(),
		Patient_Tel:     "09124214",
	}
	db.Model(&MedicalRecord{}).Create(&MedicalRecord2)

	Disease1 := Disease{
		Name:        "sore throat",
		Description: "Most sore throats are caused by infections",
	}
	db.Model(&Disease{}).Create(&Disease1)

	Disease2 := Disease{
		Name:        "headache",
		Description: "The main symptom of a headache is a pain in your head or face.",
	}
	db.Model(&Disease{}).Create(&Disease2)

	password1, err := bcrypt.GenerateFromPassword([]byte("123456"), 14)
	Nurse1 := Nurse{
		NurseName:  "nueng",
		NurseEmail: "nueng@gmail.com",
		NursePass:  string(password1),
	}
	db.Model(&Nurse{}).Create(&Nurse1)

	password2, err := bcrypt.GenerateFromPassword([]byte("123456"), 14)
	Nurse2 := Nurse{
		NurseName:  "pakapon",
		NurseEmail: "pakapon@gmail.com",
		NursePass:  string(password2),
	}
	db.Model(&Nurse{}).Create(&Nurse2)

}
