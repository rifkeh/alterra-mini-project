package controller

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io/ioutil"
	"miniproject/config"
	"miniproject/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

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

func CreateSubmissionController(c echo.Context) error{
	var submission model.Submission
	c.Bind(&submission)

	file, err := c.FormFile("files")
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "failed to get file",
		})
	}

	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "failed to open file",
		})
	}
	defer src.Close()

	fileBytes, err := ioutil.ReadAll(src)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "failed to read file",
		})
	}

	submission.File = &fileBytes

	if err := config.DB.Save(&submission).Error; err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "failed to create submission",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success create submission",
		"submission": submission,
	})
}

func GetAllSubmissionsControllerByAssignment(c echo.Context) error {
	var submissions []model.Submission
	var assignment model.Assignment
	assignmentID := c.Param("assignmentid")
	if err := config.DB.Where("class_id=?", assignmentID).Find(&submissions).Error; err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "failed to get submissions",
		})
	}

	zipBuf := new(bytes.Buffer)
	zipWriter := zip.NewWriter(zipBuf)

	for _, submission := range submissions {
		if submission.File == nil {
			continue
		}

		filename := fmt.Sprintf("submission-%d", submission.ID)
		writer, err := zipWriter.Create(filename)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": "failed to create zip entry",
			})
		}

		if _, err := writer.Write(*submission.File); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": "failed to write file to zip entry",
			})
		}
	}

	if err := zipWriter.Close(); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "failed to close zip archive",
		})
	}

	c.Response().Header().Set("Content-Type", "application/zip")

	if err := config.DB.Where("id=?", assignmentID).Find(&assignment).Error; err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "failed to get assignment",
		})
	}
	filename := fmt.Sprintf("submissions-%s.zip", assignment.Title)
	c.Response().Header().Set("Content-Disposition", "attachment; filename="+filename)

	c.Response().Write(zipBuf.Bytes())

	return nil
}
