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

func GetTeachersController(c echo.Context) error{
	var teachers []model.Teacher
	if err := config.DB.Preload("Classes").Find(&teachers).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get all teachers",
		"teachers": teachers,
	})
		
}

func CreateTeacherController(c echo.Context) error{
	var teacher model.Teacher
	var tempOTP model.Otp
	var OTP string
	c.Bind(&teacher)
	OTP = c.FormValue("OTP")
	if OTP == ""{
		tempOTP.Id = 1
		tempOTP.TeacherOTP = email.GenerateOTP()
		if err := config.DB.Save(&tempOTP).Error; err != nil {
			return echo.NewHTTPError(400, "Failed to create OTP")
		}
		if err := email.SendEmail(teacher.Name, teacher.Email, tempOTP.TeacherOTP); err != nil {
			log.Error("Failed to send account creation email:", err)
		}
		return echo.NewHTTPError(400, "See your email for OTP")

	}else{
		if err := config.DB.Find(&tempOTP).Error; err != nil {
			return echo.NewHTTPError(400, "OTP not found")
		}
		if tempOTP.TeacherOTP == OTP{
			tempOTP.Id = 1
			tempOTP.TeacherOTP = ""
			if err := config.DB.Save(&tempOTP).Error; err != nil {
				return echo.NewHTTPError(400, "Failed to create OTP")
			}
			if err := config.DB.Save(&teacher).Error; err != nil {
				return echo.NewHTTPError(400, err.Error())
			}
			return c.JSON(200, echo.Map{
				"message": "success create teacher",
				"teacher": teacher,
			})
		}else{
			return echo.NewHTTPError(400, "Wrong OTP")
		}
	}
}

func UpdateTeacherController(c echo.Context) error{
	var teacher model.Teacher
	c.Bind(&teacher)
	teacherID, _ := strconv.ParseFloat(c.Param("id"), 64)
	cookie, err := c.Cookie("TeacherSessionID")
	if err!=nil{
		return c.JSON(400, "Session expired, login again")
	}
	if middleware.ExtractTeacherIdToken(cookie.Value) == teacherID{
		if err := config.DB.Where("id = ?", teacherID).Updates(&teacher).Error; err != nil {
			return echo.NewHTTPError(400, err.Error())
		}
		return c.JSON(200, echo.Map{
			"message": "success update teacher",
			"teacher": teacher,
		})
	}
	return echo.NewHTTPError(400, "You are not authorized to update this teacher")
}

func DeleteTeacherController(c echo.Context) error{
	var teacher model.Teacher
	teacherID, _ := strconv.ParseFloat(c.Param("id"), 64)
	cookie, err := c.Cookie("TeacherSessionID")
	if err != nil{
		c.JSON(400, "Session expired, login again")
	}
	if middleware.ExtractTeacherIdToken(cookie.Value) == teacherID{
		if err := config.DB.Where("id = ?", teacherID).Delete(&teacher).Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, echo.Map{
			"message": "success delete teacher",
			"teacher": teacher,
		})
	}else{
		return c.JSON(400, "You are not authorized to delete this account")
	}
	
}

func GetTeacherController(c echo.Context)error{
	var teacher model.Teacher
	teacherID := c.Param("id")
	if err := config.DB.Where("id = ?", teacherID).Preload("Classes").Find(&teacher).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get teacher",
		"teacher": teacher,
	})
}

func LoginTeacherController(c echo.Context) error{
	teacher := model.Teacher{}
	c.Bind(&teacher)

	teachers, err := database.LoginTeacher(&teacher)

	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}
	jwtToken := teachers["Token"]
	cookie := new(http.Cookie)
	cookie.Name = "TeacherSessionID"
	cookie.Value = jwtToken
	cookie.Path = "/"
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)
	return c.JSON(200, echo.Map{
		"message": "success login teacher",
	})
}