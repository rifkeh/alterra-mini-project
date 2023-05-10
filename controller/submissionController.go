package controller

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io/ioutil"
	"miniproject/config"
	"miniproject/middleware"
	"miniproject/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetSubmissionControllerById(c echo.Context) error {
	var student model.Student
	var submission model.Submission
	var assignment model.Assignment
	submissionID := c.Param("id")
	if err := config.DB.Where("id = ?", submissionID).Find(&submission).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := config.DB.Where("id = ?", submission.StudentID).Find(&student).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := config.DB.Where("id = ?", submission.AssignmentID).Find(&assignment).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if submission.File == nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "submission does not have a file",
		})
	}
	contentType := http.DetectContentType(*submission.File)
	c.Response().Header().Set("Content-Type", contentType)
	filename := fmt.Sprintf("%s-%s-%s", student.Name, assignment.Title, submissionID)
	c.Response().Header().Set("Content-Disposition", "attachment; filename="+filename)
	c.Response().Write(*submission.File)
	return c.JSON(http.StatusOK, echo.Map{
		"message":     "success get submission",
		"submission": submission,
	})
}

func UpdateSubmissionController(c echo.Context) error {
	var submission model.Submission
	c.Bind(&submission)
	submissionID := c.Param("id")
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
	assignmentID, err := strconv.Atoi(c.Param("assignmentid"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "failed to get assignment id",
		})
	}
	cookie, err := c.Cookie("StudentSessionID")
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "failed to get cookie",
		})
	}
	submission.AssignmentID = assignmentID
	submission.StudentID = int(middleware.ExtractStudentIdToken(cookie.Value))

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
