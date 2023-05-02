package controller

import (
	"miniproject/config"
	"miniproject/lib/database"
	"miniproject/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetTeachersController(c echo.Context) error{
	var teachers []model.Teacher
	if err := config.DB.Find(&teachers).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get all teachers",
		"teachers": teachers,
	})
		
}

func CreateTeacherController(c echo.Context) error{
	var teacher model.Teacher
	c.Bind(&teacher)

	if err:=config.DB.Save(&teacher).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success create teacher",
		"teacher": teacher,
	})
}

func UpdateTeacherController(c echo.Context) error{
	var teacher model.Teacher
	c.Bind(&teacher)
	teacherID := c.Param("id")
	if err := config.DB.Where("id = ?", teacherID).Updates(&teacher).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success update teacher",
		"teacher": teacher,
	})
}

func DeleteTeacherController(c echo.Context) error{
	var teacher model.Teacher
	teacherID := c.Param("id")
	if err := config.DB.Where("id = ?", teacherID).Delete(&teacher).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success delete teacher",
		"teacher": teacher,
	})
}

func GetTeacherController(c echo.Context)error{
	var teacher model.Teacher
	teacherID := c.Param("id")
	if err := config.DB.Where("id = ?", teacherID).Find(&teacher).Error; err != nil {
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
	return c.JSON(200, echo.Map{
		"message": "success login teacher",
		"teacher": teachers,
	})
}