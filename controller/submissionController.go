package controller

import (
	"miniproject/config"
	"miniproject/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetSubmissionsController(c echo.Context) error {
	var submissions []model.Submission
	if err := config.DB.Find(&submissions).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message":     "success get all submissions",
		"submissions": submissions,
	})
}

func CreateSubmissionController(c echo.Context) error {
	var submission model.Submission
	c.Bind(&submission)

	if err := config.DB.Save(&submission).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message":     "success create submission",
		"submission": submission,
	})
}

func UpdateSubmissionController(c echo.Context) error {
	var submission model.Submission
	c.Bind(&submission)
	submissionID := c.Param("id")
	if err := config.DB.Where("id = ?", submissionID).Updates(&submission).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message":     "success update submission",
		"submission": submission,
	})
}

func DeleteSubmissionController(c echo.Context) error {
	var submission model.Submission
	submissionID := c.Param("id")
	if err := config.DB.Where("id = ?", submissionID).Delete(&submission).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message":     "success delete submission",
		"submission": submission,
	})
}

func GetSubmissionController(c echo.Context) error {
	var submission model.Submission
	submissionID := c.Param("id")
	if err := config.DB.Where("id = ?", submissionID).First(&submission).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message":     "success get submission",
		"submission": submission,
	})
}
