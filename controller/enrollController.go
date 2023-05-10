package controller

import (
	"miniproject/config"
	"miniproject/middleware"
	"miniproject/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetEnrollsController(c echo.Context) error {
	var enrolls []model.Enrollment
	if err := config.DB.Find(&enrolls).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get all enrolls",
		"enrolls": enrolls,
	})
}

func CreateEnrollController(c echo.Context) error {
	var enroll model.Enrollment
	var class model.Class
	c.Bind(&enroll)
	classID, _ := strconv.Atoi(c.Param("classid"))
	cookie, err := c.Cookie("StudentSessionID")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	studentID := int (middleware.ExtractStudentIdToken(cookie.Value))
	enroll.StudentID = studentID
	enroll.ClassID = classID
	if err := config.DB.Where("id = ?", classID).First(&class).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if class.Password != enroll.Password{
		return echo.NewHTTPError(http.StatusBadRequest, "wrong password")
	}
	if err := config.DB.Save(&enroll).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success create enroll",
		"enroll":  enroll,
	})
}

func DeleteEnrollController(c echo.Context) error {
	var enroll model.Enrollment
	enrollID := c.Param("id")
	if err := config.DB.Where("id = ?", enrollID).Delete(&enroll).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success delete enroll",
		"enroll":  enroll,
	})
}

func GetEnrollController(c echo.Context) error {
	var enroll model.Enrollment
	enrollID := c.Param("id")
	if err := config.DB.Where("id = ?", enrollID).Find(&enroll).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get enroll",
		"enroll":  enroll,
	})
}