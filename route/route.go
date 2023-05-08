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
	e.POST("/class", controller.CreateClassController)
	e.PUT("/class/:id", controller.UpdateClassController)
	e.DELETE("/class/:id", controller.DeleteClassController)
	e.GET("/class/:id", controller.GetClassController)
	// Student Routes
	s.GET("/students", controller.GetStudentsController)
	t.GET("/students", controller.GetStudentsController)
	t.GET("/student/:id", controller.GetStudentController)
	e.POST("/student/register", controller.CreateStudentController)
	s.PUT("/student/:id", controller.UpdateStudentController)
	s.DELETE("/student/:id", controller.DeleteStudentController)
	s.GET("/student/:id", controller.GetStudentController)
	e.POST("/student/login", controller.LoginStudentController)
	// Enrollment Routes
	s.GET("/enrollments", controller.GetEnrollController)
	s.POST("/enrollment", controller.CreateEnrollController)
	s.DELETE("/enrollment/:id", controller.DeleteEnrollController)
	s.GET("/enrollment/:id", controller.GetEnrollController)
	t.DELETE("/enrollment/:id", controller.DeleteEnrollController)
	t.GET("/enrollments", controller.GetEnrollController)
	t.GET("/enrollment/:id", controller.GetEnrollController)
	// Assignment Routes
	t.GET("/assignments", controller.GetAssignmentsController)
	t.POST("/assignment", controller.CreateAssignmentController)
	t.PUT("/assignment/:id", controller.UpdateAssignmentController)
	t.DELETE("/assignment/:id", controller.DeleteAssignmentController)
	t.GET("/assignment/:id", controller.GetAssignmentController)
	s.GET("/assignments", controller.GetAssignmentsController)
	s.GET("/assignment/:id", controller.GetAssignmentController)
	// Submission Routes
	s.DELETE("/assignment/:assignmentid/submission/:id", controller.DeleteSubmissionController)
	s.POST("/assignment/:assigmentid/submission", controller.CreateSubmissionController)
	s.GET("/assignment/:assignmentid/submission", controller.GetAllSubmissionsControllerByAssignment)
	// Material Routes
	t.GET("/materials", controller.GetMaterialsController)
	t.POST("/material", controller.CreateMaterialController)
	t.PUT("/material/:id", controller.UpdateMaterialController)
	t.DELETE("/material/:id", controller.DeleteMaterialController)
	t.GET("/material/:id", controller.GetMaterialController)
	s.GET("/materials", controller.GetMaterialsController)
	s.GET("/material/:id", controller.GetMaterialController)
	return e
}