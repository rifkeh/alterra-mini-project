package controller

import (
	"miniproject/config"
	"miniproject/middleware"
	"miniproject/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetClassesController(c echo.Context) error {
	var classes []model.Class
	if err := config.DB.Limit(5).Find(&classes).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get all classes",
		"classes": classes,
	})
}

func CreateClassController(c echo.Context) error {
	var class model.Class
	c.Bind(&class)
	cookie, err := c.Cookie("TeacherSessionID")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	jwtToken := cookie.Value
	class.TeacherID = int(middleware.ExtractTeacherIdToken(jwtToken))
	if err := config.DB.Save(&class).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success create class",
		"class":   class,
	})
}

func UpdateClassController(c echo.Context) error {
	var class model.Class
	cookie, err := c.Cookie("TeacherSessionID")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	classID := c.Param("id")
	c.Bind(&class)
	class.TeacherID = int(middleware.ExtractTeacherIdToken(cookie.Value))
	if err := config.DB.Where("teacher_id = ? AND id = ?", int(middleware.ExtractTeacherIdToken(cookie.Value)),classID).Updates(&class).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Class does not exist")
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success update class",
		"class":   class,
	})
}

func DeleteClassController(c echo.Context) error {
	var class model.Class
	classID := c.Param("id")
	cookie, err := c.Cookie("TeacherSessionID")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := config.DB.Where("teacher_id=? AND id = ?", int(middleware.ExtractTeacherIdToken(cookie.Value)) ,classID).Delete(&class).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success delete class",
		"class":   class,
	})
}

func GetClassController(c echo.Context) error {
	var class model.Class
	classID := c.Param("id")
	if err := config.DB.Where("id = ?", classID).Find(&class).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get class",
		"class":   class,
	})
}