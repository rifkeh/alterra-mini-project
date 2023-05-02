package controller

import (
	"miniproject/config"
	"miniproject/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetClassesController(c echo.Context) error {
	var classes []model.Class
	if err := config.DB.Find(&classes).Error; err != nil {
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
	c.Bind(&class)
	classID := c.Param("id")
	if err := config.DB.Where("id = ?", classID).Updates(&class).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success update class",
		"class":   class,
	})
}

func DeleteClassController(c echo.Context) error {
	var class model.Class
	classID := c.Param("id")
	if err := config.DB.Where("id = ?", classID).Delete(&class).Error; err != nil {
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