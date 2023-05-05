package model

import (
	"github.com/jinzhu/gorm"
)

type Student struct {
	gorm.Model
	Name       string	`json:"name" form:"name"`
	Email      string	`json:"email" form:"email"`
	Password   string	`json:"password" form:"password"`
	Enrollment []Enrollment `gorm:"foreignKey:StudentID"`
	Submission []Submission `gorm:"foreignKey:StudentID"`
	// LastLogin  time.Time	`json:"last_login"`
}

type Teacher struct {
	gorm.Model
	Name       string	`json:"name" form:"name"`
	Email      string	`json:"email" form:"email"`
	Password   string	`json:"password" form:"password"`
	Classes    []Class `gorm:"foreignKey:TeacherID"`
	// LastLogin  time.Time	`json:"last_login"`
}

type Enrollment struct {
	gorm.Model
	StudentID int	`json:"student_id" form:"student_id"`
	ClassID   int	`json:"class_id" form:"class_id"`
}

type Class struct {
	gorm.Model
	TeacherID   int	`json:"teacher_id" form:"teacher_id"`
	Name        string	`json:"name" form:"name"`
	Description string	`json:"description" form:"description"`
	Assignment  []Assignment `gorm:"foreignKey:ClassID"`
	Material    []Material`gorm:"foreignKey:ClassID"`
}

type Assignment struct {
	gorm.Model
	ClassID     int	`json:"class_id" form:"class_id"`
	Description string	`json:"description" form:"description"`
	Deadline    string	`json:"deadline" form:"deadline"`
}

type Material struct {
	gorm.Model
	ClassID     int	`json:"class_id" form:"class_id"`
	Description string	`json:"description" form:"description"`
}

type Submission struct {
	gorm.Model
	AssignmentID int	`json:"assignment_id" form:"assignment_id"`
	StudentID    int	`json:"student_id" form:"student_id"`
	File         []byte	`json:"file" form:"file"`
}

type Otp struct {
	Id	int ``
	StudentOTP string `json:"student_otp" form:"student_otp"`
	TeacherOTP string `json:"teacher_otp" form:"teacher_otp"`
}