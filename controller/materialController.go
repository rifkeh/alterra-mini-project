package controller

import (
	"fmt"
	"miniproject/config"
	"miniproject/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetMaterialsController(c echo.Context) error {
	var materials []model.Material
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
		"material": material,
	})
}

func GetMaterialController(c echo.Context) error {
	var material model.Material
	materialID := c.Param("id")
	if err := config.DB.Where("id = ?", materialID).Find(&material).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message":   "success get material",
		"material": material,
	})
}
