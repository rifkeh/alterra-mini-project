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
	e.GET("/teacher/:id", controller.GetTeacherController)
	e.GET("/classes", controller.GetClassesController)
	e.POST("/class", controller.CreateClassController)
	e.PUT("/class/:id", controller.UpdateClassController)
	e.DELETE("/class/:id", controller.DeleteClassController)
	e.GET("/class/:id", controller.GetClassController)
	e.GET("/students", controller.GetStudentsController)
	e.POST("/student", controller.CreateStudentController)
	e.PUT("/student/:id", controller.UpdateStudentController)
	e.DELETE("/student/:id", controller.DeleteStudentController)
	e.GET("/student/:id", controller.GetStudentController)
	e.GET("/enrollments", controller.GetEnrollController)
	e.POST("/enrollment", controller.CreateEnrollController)
	e.PUT("/enrollment/:id", controller.UpdateEnrollController)
	e.DELETE("/enrollment/:id", controller.DeleteEnrollController)
	e.GET("/enrollment/:id", controller.GetEnrollController)
	e.GET("/assignments", controller.GetAssignmentsController)
	e.POST("/assignment", controller.CreateAssignmentController)
	e.PUT("/assignment/:id", controller.UpdateAssignmentController)
	e.DELETE("/assignment/:id", controller.DeleteAssignmentController)
	e.GET("/assignment/:id", controller.GetAssignmentController)
	e.GET("/submissions", controller.GetSubmissionsController)
	e.POST("/submission", controller.CreateSubmissionController)
	e.PUT("/submission/:id", controller.UpdateSubmissionController)
	e.DELETE("/submission/:id", controller.DeleteSubmissionController)
	e.GET("/submission/:id", controller.GetSubmissionController)
	e.GET("/materials", controller.GetMaterialsController)
	e.POST("/material", controller.CreateMaterialController)
	e.PUT("/material/:id", controller.UpdateMaterialController)
	e.DELETE("/material/:id", controller.DeleteMaterialController)
	e.GET("/material/:id", controller.GetMaterialController)
	return e
}