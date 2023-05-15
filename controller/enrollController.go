package controller

import (
	"fmt"
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

func GetEnrollsController(c echo.Context) error {
	var enrolls []model.Enrollment
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
	classID := c.Param("classid")
	if err := config.DB.Where("class_id=?", classID).Find(&enrolls).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get all enrolls",
		"enrolls": enrolls,
	})
}

func CreateEnrollController(c echo.Context) error {
	var enroll model.Enrollment
	var class model.Class
	c.Bind(&enroll)
	classID, _ := strconv.Atoi(c.Param("classid"))
	cookie, err := c.Cookie("StudentSessionID")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	studentID := int (middleware.ExtractStudentIdToken(cookie.Value))
	enroll.StudentID = studentID
	enroll.ClassID = classID
	if err := config.DB.Where("id = ?", classID).First(&class).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if class.Password != enroll.Password{
		return echo.NewHTTPError(http.StatusBadRequest, "wrong password")
	}
	if err := config.DB.Save(&enroll).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success create enroll",
		"enroll":  enroll,
	})
}

func DeleteEnrollController(c echo.Context) error {
	var enroll model.Enrollment
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
	enrollID := c.Param("id")
	classID := c.Param("classid")
	if err := config.DB.Where("id = ? AND class_id=?", enrollID, classID).Find(&enroll).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get enroll",
		"enroll":  enroll,
	})
}