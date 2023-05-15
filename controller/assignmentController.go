package controller

import (
	"fmt"
	"io/ioutil"
	"miniproject/config"
	"miniproject/constant"
	"miniproject/middleware"
	"miniproject/model"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func GetAssignmentsControllerByClass(c echo.Context) error {
	var assignments []model.Assignment
	classID, _ := strconv.Atoi(c.Param("classid"))
	if err := config.DB.Where("class_id=?", classID).Find(&assignments).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message":    "success get all assignments",
		"assignments": assignments,
	})
}

func CreateAssignmentController(c echo.Context) error {
	var assignment model.Assignment
	var class model.Class
	c.Bind(&assignment)
	classID := c.Param("classid")
	assignment.ClassID, _ = strconv.Atoi(classID)
	if err := config.DB.Where("id = ?", classID).First(&class).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	cookie, err := c.Cookie("TeacherSessionID")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if class.TeacherID != int(middleware.ExtractTeacherIdToken(cookie.Value)) {
		return echo.NewHTTPError(http.StatusBadRequest, "you are not the teacher of this class")
	}
	file, err := c.FormFile("file")
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

	assignment.File = &fileBytes

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
	classID := c.Param("classid")

	authHeader := c.Request().Header.Get("Authorization")
	tokenString := strings.Split(authHeader, " ")[1]
	if tokenString == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
	}

	// Parse and validate the JWT token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		if token.Claims.(jwt.MapClaims)["teacherID"] != nil {
			return []byte(constant.TEACHER_JWT), nil
		} else if token.Claims.(jwt.MapClaims)["studentID"] != nil {
			return []byte(constant.STUDENT_JWT), nil
		}

		return nil, fmt.Errorf("invalid token")
	})

	if err != nil || !token.Valid {
		return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
	}

	if err := config.DB.Where("id = ? AND class_id = ?", assignmentID, classID).First(&assignment).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	contentType := http.DetectContentType(*assignment.File)
	c.Response().Header().Set("Content-Type", contentType)
	filename := fmt.Sprintf("submission-%s", assignmentID)
	c.Response().Header().Set("Content-Disposition", "attachment; filename="+filename)
	c.Response().Write(*assignment.File)

	return c.JSON(http.StatusOK, echo.Map{
		"message":    "success get assignment",
		"assignment": assignment,
	})
}

