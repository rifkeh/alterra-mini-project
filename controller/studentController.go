package controller

import (
	"miniproject/config"
	"miniproject/lib/database"
	"miniproject/lib/email"
	"miniproject/middleware"
	"miniproject/model"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func GetStudentsController(c echo.Context) error {
	var students []model.Student
	if err := config.DB.Preload("Enrollment").Find(&students).Error; err != nil {
		return echo.NewHTTPError(400, err.Error())
	}
	studentclass := make([]model.StudentClass, len(students))
	
	for s, i := range students{
		studentclass[s].Student = append(studentclass[s].Student, i)
		for _, j := range i.Enrollment{
			var classes model.Class
			if err:=config.DB.Where("id = ?", j.ClassID).Find(&classes).Error; err != nil{
				return echo.NewHTTPError(400, "Failed to get class")
			}
			studentclass[s].Class = append(studentclass[s].Class, classes)
		}
	}
	
	return c.JSON(200, echo.Map{
		"message": "success get all students",
		"students": studentclass,
	})
}

func CreateStudentController(c echo.Context) error {
	var student model.Student
	var tempOTP model.Otp
	var OTP string
	c.Bind(&student)
	OTP = c.FormValue("OTP")
	if OTP == ""{
		tempOTP.Id = 1
		tempOTP.StudentOTP = email.GenerateOTP()
		if err := config.DB.Save(&tempOTP).Error; err != nil {
			return echo.NewHTTPError(400, "Failed to create OTP")
		}
		if err := email.SendEmail(student.Name, student.Email, tempOTP.StudentOTP); err != nil {
			log.Error("Failed to send account creation email:", err)
		}
		return echo.NewHTTPError(400, "See your email for OTP")

	}else{
		if err := config.DB.Find(&tempOTP).Error; err != nil {
			return echo.NewHTTPError(400, "OTP not found")
		}
		if tempOTP.StudentOTP == OTP{
			tempOTP.Id = 1
			tempOTP.StudentOTP = ""
			if err := config.DB.Save(&tempOTP).Error; err != nil {
				return echo.NewHTTPError(400, "Failed to create OTP")
			}
			if err := config.DB.Save(&student).Error; err != nil {
				return echo.NewHTTPError(400, err.Error())
			}
			return c.JSON(200, echo.Map{
				"message": "success create student",
				"student": student,
			})
		}else{
			return echo.NewHTTPError(400, "Wrong OTP")
		}
	}
}
	

func UpdateStudentController(c echo.Context) error {
	var student model.Student
	c.Bind(&student)
	studentID, _ := strconv.ParseFloat(c.Param("id"), 64)
	cookie, err := c.Cookie("StudentSessionID")
	if err != nil{
		return c.JSON(400, "Session expired, login again")
	}
	if middleware.ExtractStudentIdToken(cookie.Value) == studentID{
		if err := config.DB.Where("id = ?", studentID).Updates(&student).Error; err != nil {
			return echo.NewHTTPError(400, err.Error())
		}
		return c.JSON(200, echo.Map{
			"message": "success update student",
			"student": student,
		})
	}
	return echo.NewHTTPError(400, "You are not authorized to update this student")
	
}

func DeleteStudentController(c echo.Context) error {
	var student model.Student
	studentID, _ := strconv.ParseFloat(c.Param("id"), 64)
	cookie, err := c.Cookie("StudentSessionID")
	if err != nil{
		return c.JSON(400, "Session expired, login again")
	}
	if middleware.ExtractStudentIdToken(cookie.Value) == studentID{
		if err := config.DB.Where("id = ?", studentID).Delete(&student).Error; err != nil {
			return echo.NewHTTPError(400, err.Error())
		}
		return c.JSON(200, echo.Map{
			"message": "success delete student",
			"student": student,
		})
	}
	return echo.NewHTTPError(400, "You are not authorized to delete this student")
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
	jwtToken := students["Token"]
	cookie := new(http.Cookie)
	cookie.Name = "StudentSessionID"
	cookie.Value = jwtToken
	cookie.Path = "/"
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)
	return c.JSON(200, echo.Map{
		"message": "success login student",
	})
}