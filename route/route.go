package route

import (
	"miniproject/controller"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo{
	e := echo.New()
	e.GET("/teachers", controller.GetTeachersController)
	e.POST("/teacher", controller.CreateTeacherController)
	e.PUT("/teacher/:id", controller.UpdateTeacherController)
	e.DELETE("/teacher/:id", controller.DeleteTeacherController)
	return e
}