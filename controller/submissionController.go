package controller

import (
	"fmt"
	"io/ioutil"
	"miniproject/config"
	"miniproject/middleware"
	"miniproject/model"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func GetAllSubmissionsControllerById(c echo.Context) error {
	var submissions []model.Submission
	var student model.Student
	AssignmentID := c.Param("id")
	if err := config.DB.Where("assignment_id=?", AssignmentID).Find(&submissions).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	for _, file := range submissions{
		if err := config.DB.Where("id=?", file.StudentID).Find(&student).Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		contentType := http.DetectContentType(file.File)
		contentDisposition := fmt.Sprintf("attachment; filename=submission-%s-%s", AssignmentID, student.Name)
		c.Response().Header().Set("Content-Type", contentType)
		c.Response().Header().Set("Content-Disposition", contentDisposition)
	}
	
	return c.JSON(http.StatusOK, echo.Map{
		"message":     "success get all submissions",
		"submissions": submissions,
	})
}

func CreateSubmissionController(c echo.Context) error {
	var submission model.Submission
	c.Bind(&submission)
	if middleware.ExtractStudentIdToken(strings.Replace(c.Request().Header.Get("Authorization"), "Bearer ", "", -1)) == float64(submission.StudentID){
		file, err := c.FormFile("file")
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		src, err := file.Open()
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		data , err := ioutil.ReadAll(src)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		
		defer src.Close()
		submission.File = data
	
		if err := config.DB.Save(&submission).Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, echo.Map{
			"message":     "success create submission",
			"submission": submission,
		})
	}
	return echo.NewHTTPError(http.StatusBadRequest, "You are not authorized to access this submission")
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
	var student model.Student
	submissionID := c.Param("id")
	if err := config.DB.Where("id = ?", submissionID).First(&submission).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := config.DB.Where("id = ?", submission.StudentID).First(&student).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	contentType := http.DetectContentType(submission.File)
    contentDisposition := fmt.Sprintf("attachment; filename=submission-%s-%s", submissionID, student.Name)
    c.Response().Header().Set("Content-Type", contentType)
    c.Response().Header().Set("Content-Disposition", contentDisposition)

    if _, err := c.Response().Write(submission.File); err != nil {
        return err
    }
	return c.JSON(http.StatusOK, echo.Map{
		"message":     "success get submission",
		"submission": submission,
	})
}
