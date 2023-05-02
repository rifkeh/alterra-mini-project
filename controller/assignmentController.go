package controller

import (
	"miniproject/config"
	"miniproject/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAssignmentsController(c echo.Context) error {
	var assignments []model.Assignment
	if err := config.DB.Find(&assignments).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message":    "success get all assignments",
		"assignments": assignments,
	})
}

func CreateAssignmentController(c echo.Context) error {
	var assignment model.Assignment
	c.Bind(&assignment)

	if err := config.DB.Save(&assignment).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message":    "success create assignment",
		"assignment": assignment,
	})
}

func UpdateAssignmentController(c echo.Context) error {
	var assignment model.Assignment
	c.Bind(&assignment)
	assignmentID := c.Param("id")
	if err := config.DB.Where("id = ?", assignmentID).Updates(&assignment).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message":    "success update assignment",
		"assignment": assignment,
	})
}

func DeleteAssignmentController(c echo.Context) error {
	var assignment model.Assignment
	assignmentID := c.Param("id")
	if err := config.DB.Where("id = ?", assignmentID).Delete(&assignment).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message":    "success delete assignment",
		"assignment": assignment,
	})
}

func GetAssignmentController(c echo.Context) error {
	var assignment model.Assignment
	assignmentID := c.Param("id")
	if err := config.DB.Where("id = ?", assignmentID).First(&assignment).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message":    "success get assignment",
		"assignment": assignment,
	})
}
