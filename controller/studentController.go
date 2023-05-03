package controller

import (
	"miniproject/config"
	"miniproject/lib/database"
	"miniproject/model"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func GetStudentsController(c echo.Context) error {
	var students []model.Student
	if err := config.DB.Find(&students).Error; err != nil {
		return echo.NewHTTPError(400, err.Error())
	}
	return c.JSON(200, echo.Map{
		"message": "success get all students",
		"students": students,
	})
}

func CreateStudentController(c echo.Context) error {
	var student model.Student
	var OTP string
	c.Bind(&student)
	if err := database.SendEmail(student.Email); err != nil {
		log.Error("Failed to send account creation email:", err)
	}
	if OTP == "test"{
		if err := config.DB.Save(&student).Error; err != nil {
			return echo.NewHTTPError(400, err.Error())
		}
		return c.JSON(200, echo.Map{
			"message": "success create student",
			"student": student,
		})
	} else {
		return echo.NewHTTPError(400, "OTP is wrong")
	}
}
	

func UpdateStudentController(c echo.Context) error {
	var student model.Student
	c.Bind(&student)
	studentID := c.Param("id")
	if err := config.DB.Where("id = ?", studentID).Updates(&student).Error; err != nil {
		return echo.NewHTTPError(400, err.Error())
	}
	return c.JSON(200, echo.Map{
		"message": "success update student",
		"student": student,
	})
}

func DeleteStudentController(c echo.Context) error {
	var student model.Student
	studentID := c.Param("id")
	if err := config.DB.Where("id = ?", studentID).Delete(&student).Error; err != nil {
		return echo.NewHTTPError(400, err.Error())
	}
	return c.JSON(200, echo.Map{
		"message": "success delete student",
		"student": student,
	})
}

func GetStudentController(c echo.Context) error {
	var student model.Student
	studentID := c.Param("id")
	if err := config.DB.Where("id = ?", studentID).First(&student).Error; err != nil {
		return echo.NewHTTPError(400, err.Error())
	}
	return c.JSON(200, echo.Map{
		"message": "success get student",
		"student": student,
	})
}

func LoginStudentController(c echo.Context) error{
	student := model.Student{}
	c.Bind(&student)

	students, err := database.LoginStudent(&student)

	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}
	return c.JSON(200, echo.Map{
		"message": "success login student",
		"student": students,
	})
}