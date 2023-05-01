package controller

import (
	"miniproject/config"
	"miniproject/model"
	"net/http"

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
	c.Bind(&enroll)

	if err := config.DB.Save(&enroll).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success create enroll",
		"enroll":  enroll,
	})
}

func UpdateEnrollController(c echo.Context) error {
	var enroll model.Enrollment
	c.Bind(&enroll)
	enrollID := c.Param("id")
	if err := config.DB.Where("id = ?", enrollID).Updates(&enroll).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success update enroll",
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