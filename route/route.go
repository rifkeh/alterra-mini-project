package route

import (
	"miniproject/constant"
	"miniproject/controller"
	"miniproject/middleware"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo{
	e := echo.New()
	t := e.Group("")
	s := e.Group("")
	s.Use(echojwt.JWT([]byte(constant.STUDENT_JWT)))
	t.Use(echojwt.JWT([]byte(constant.TEACHER_JWT)))
	middleware.LogMiddleware(e)
	// Teacher Routes
	t.GET("/teachers", controller.GetTeachersController)
	e.POST("/teacher/register", controller.CreateTeacherController)
	t.PUT("/teacher/:id", controller.UpdateTeacherController)
	t.DELETE("/teacher/:id", controller.DeleteTeacherController)
	t.GET("/teacher/:id", controller.GetTeacherController)
	e.POST("/teacher/login", controller.LoginTeacherController)
	// Class Routes
	e.GET("/classes", controller.GetClassesController)
	t.POST("/class", controller.CreateClassController)
	t.PUT("/class/:id", controller.UpdateClassController)
	t.DELETE("/class/:id", controller.DeleteClassController)
	e.GET("/class/:id", controller.GetClassController)
	// Student Routes
	t.GET("/students", controller.GetStudentsController)
	s.GET("/students", controller.GetStudentsController)
	t.GET("/student/:id", controller.GetStudentController)
	e.POST("/student/register", controller.CreateStudentController)
	s.PUT("/student/:id", controller.UpdateStudentController)
	s.DELETE("/student/:id", controller.DeleteStudentController)
	s.GET("/student/:id", controller.GetStudentController)
	e.POST("/student/login", controller.LoginStudentController)
	// Enrollment Routes
	s.GET("/class/:classid/enrollments", controller.GetEnrollController)
	s.POST("/class/:classid/enrollment", controller.CreateEnrollController)
	s.DELETE("/class/:classid/enrollment/:id", controller.DeleteEnrollController)
	s.GET("/class/:classid/enrollment/:id", controller.GetEnrollController)
	t.DELETE("/class/:classid/enrollment/:id", controller.DeleteEnrollController)
	t.GET("/class/:classid/enrollments", controller.GetEnrollController)
	t.GET("/class/:classid/enrollment/:id", controller.GetEnrollController)
	// Assignment Routes
	t.GET("/class/:classid/assignments", controller.GetAssignmentsControllerByClass)
	t.POST("/class/:classid/assignment", controller.CreateAssignmentController)
	t.PUT("/class/:classid/assignment/:id", controller.UpdateAssignmentController)
	t.DELETE("/class/:classid/assignment/:id", controller.DeleteAssignmentController)
	t.GET("/class/:classid/assignment/:id", controller.GetAssignmentController)
	s.GET("/class/:classid/assignments", controller.GetAssignmentsControllerByClass)
	s.GET("/class/:classid/assignment/:id", controller.GetAssignmentController)
	// Submission Routes
	s.DELETE("/assignment/:assignmentid/submission/:id", controller.DeleteSubmissionController)
	s.POST("/assignment/:assignmentid/submission", controller.CreateSubmissionController)
	s.GET("/assignment/:assignmentid/submission", controller.GetAllSubmissionsControllerByAssignment)
	s.GET("/assignment/:assignmentid/submission/:id", controller.GetSubmissionControllerById)
	// Material Routes
	t.GET("/class/:classid/materials", controller.GetMaterialsController)
	t.POST("/class/:classid/material", controller.CreateMaterialController)
	t.PUT("/class/:classid/material/:id", controller.UpdateMaterialController)
	t.DELETE("/class/:classid/material/:id", controller.DeleteMaterialController)
	t.GET("/class/:classid/material/:id", controller.GetMaterialController)
	s.GET("/class/:classid/materials", controller.GetMaterialsController)
	s.GET("/class/:classid/material/:id", controller.GetMaterialController)
	return e
}