package controller

import (
	"fmt"
	"miniproject/config"
	"miniproject/constant"
	"miniproject/model"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func GetMaterialsController(c echo.Context) error {
	var materials []model.Material
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
	if err := config.DB.Limit(5).Find(&materials).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message":   "success get all materials",
		"materials": materials,
	})
}

func GetMaterialsControllerByClass(c echo.Context) error {
	var materials []model.Material
	classID := c.Param("classid")
	if err := config.DB.Where("class_id=?", classID).Find(&materials).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message":   fmt.Sprintf("success get all materials from class %s", classID),
		"materials": materials,
	})
}

func CreateMaterialController(c echo.Context) error {
	var material model.Material
	c.Bind(&material)
	classID, _ := strconv.Atoi(c.Param("classid"))
	material.ClassID = classID
	if err := config.DB.Save(&material).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message":   "success create material",
		"material": material,
	})
}

func UpdateMaterialController(c echo.Context) error {
	var material model.Material
	c.Bind(&material)
	materialID := c.Param("id")
	classID := c.Param("classid")
	if err := config.DB.Where("id = ? AND class_id=?", materialID, classID).Updates(&material).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message":   "success update material",
		"material": material,
	})
}

func DeleteMaterialController(c echo.Context) error {
	var material model.Material
	materialID := c.Param("id")
	if err := config.DB.Where("id = ?", materialID).Delete(&material).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message":   "success delete material",
	})
}

func GetMaterialController(c echo.Context) error {
	var material model.Material
	materialID := c.Param("id")
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
	if err := config.DB.Where("id = ?", materialID).Find(&material).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message":   "success get material",
		"material": material,
	})
}
