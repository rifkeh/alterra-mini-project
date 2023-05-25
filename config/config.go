package config

import (
	"fmt"

	"miniproject/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	config := map[string]string{
		"DB_Username": "root",
		"DB_Password": "rifkhi",
		"DB_Port":     "3306",
		"DB_Host":     "db_mysql",
		"DB_Name":     "LMS",
	}
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", config["DB_Username"], config["DB_Password"], config["DB_Host"], config["DB_Port"], config["DB_Name"])
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	InitMigrate()
}

func InitMigrate() {
	DB.AutoMigrate(&model.Student{})
	DB.AutoMigrate(&model.Teacher{})
	DB.AutoMigrate(&model.Class{})
	DB.AutoMigrate(&model.Enrollment{})
	DB.AutoMigrate(&model.Assignment{})
	DB.AutoMigrate(&model.Material{})
	DB.AutoMigrate(&model.Submission{})
	DB.AutoMigrate(&model.Otp{})
}
