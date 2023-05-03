package controller

import (
	"fmt"
	"io/ioutil"
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
	// filename := "submission-" + submissionID + ".png"
	// if err := ioutil.WriteFile(filename, submission.File, 0644); err != nil {
	// 	return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	// }
	// defer func() {
	// 	if err := os.Remove(filename); err != nil {
	// 		panic(err)
	// 	}
	// }()
	contentType := http.DetectContentType(submission.File)
    contentDisposition := fmt.Sprintf("attachment; filename=%s", "test")
    c.Response().Header().Set("Content-Type", contentType)
    c.Response().Header().Set("Content-Disposition", contentDisposition)

    // Write the file data to the response body
    if _, err := c.Response().Write(submission.File); err != nil {
        return err
    }
	return c.JSON(http.StatusOK, echo.Map{
		"message":     "success get submission",
		"submission": submission,
	})
}
