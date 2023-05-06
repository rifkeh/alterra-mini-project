package model

import (
	"github.com/jinzhu/gorm"
)

type Student struct {
	gorm.Model
	Name       string	`json:"name" form:"name" gorm:"unique"`
	Email      string	`json:"email" form:"email" gorm:"unique"`
	Password   string	`json:"password" form:"password" gorm:"notnull"`
	Enrollment []Enrollment `gorm:"foreignKey:StudentID" json:"enrollment"`
	Submission []Submission `gorm:"foreignKey:StudentID" json:"submission"`
}

type Teacher struct {
	gorm.Model
	Name       string	`json:"name" form:"name"`
	Email      string	`json:"email" form:"email"`
	Password   string	`json:"password" form:"password"`
	Classes    []Class `gorm:"foreignKey:TeacherID"`
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
	Password	string	`json:"password" form:"password"`
	Assignment  []Assignment `gorm:"foreignKey:ClassID"`
	Material    []Material`gorm:"foreignKey:ClassID"`
	Enrollment []Enrollment `gorm:"foreignKey:ClassID"`
}

type Assignment struct {
	gorm.Model
	ClassID     int	`json:"class_id" form:"class_id"`
	Title	   string	`json:"title" form:"title" gorm:"unique;not null"`
	Description string	`json:"description" form:"description" gorm:"not null"`
	Deadline    string	`json:"deadline" form:"deadline" gorm:"not null"`
}

type Material struct {
	gorm.Model
	ClassID     int	`json:"class_id" form:"class_id"`
	Title	   string	`json:"title" form:"title"`
	Description string	`json:"description" form:"description"`
	File		*[]byte	`json:"file" form:"file"`
}

type Submission struct {
	gorm.Model
	AssignmentID int	`json:"assignment_id" form:"assignment_id" gorm:"unique;not null"`
	StudentID    int	`json:"student_id" form:"student_id" gorm:"unique;not null"`
	File         *[]byte	`json:"file" form:"file" gorm:"not null"`
	Comment	  	string	`json:"comment" form:"comment"`
	Grade 	   	int	`json:"grade" form:"grade"`
}

type Otp struct {
	Id	int ``
	StudentOTP string `json:"student_otp" form:"student_otp"`
	TeacherOTP string `json:"teacher_otp" form:"teacher_otp"`
}