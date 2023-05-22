package model

import (
	"github.com/jinzhu/gorm"
)

type StudentClass struct {
	Student []Student `gorm:"many2many:student_classes;"`
	Class []Class `gorm:"many2many:student_classes;"`
}

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
	Name       string	`json:"name" form:"name" gorm:"unique"`
	Email      string	`json:"email" form:"email" gorm:"unique"`
	Password   string	`json:"password" form:"password" gorm:"notnull"`
	Classes    []Class `gorm:"foreignKey:TeacherID"`
}

type Enrollment struct {
	gorm.Model
	StudentID int	`json:"student_id" form:"student_id"`
	ClassID   int	`json:"class_id" form:"class_id"`
	Password string	`json:"password" form:"password"`
}

type Class struct {
	gorm.Model
	TeacherID   int	`json:"teacher_id" form:"teacher_id" gorm:"not null"`
	Name        string	`json:"name" form:"name" gorm:"unique;not null"`
	Description string	`json:"description" form:"description" gorm:"not null"`
	Password	string	`json:"password" form:"password"`
	Assignment  []Assignment `gorm:"foreignKey:ClassID"`
	Material    []Material`gorm:"foreignKey:ClassID"`
	Enrollment []Enrollment `gorm:"foreignKey:ClassID"`
}

type Assignment struct {
	gorm.Model
	ClassID     int	`json:"class_id" form:"class_id"`
	Title	   string	`json:"title" form:"title" gorm:"not null"`
	Description string	`json:"description" form:"description" gorm:"not null"`
	File		*[]byte	`json:"file" form:"file"`
	Deadline    string	`json:"deadline" form:"deadline" gorm:"not null"`
	Submission Submission `gorm:"foreignKey:AssignmentID"`
}

type Material struct {
	gorm.Model
	ClassID     int	`json:"class_id" form:"class_id" gorm:"not null"`
	Title	   string	`json:"title" form:"title" gorm:"unique;not null"`
	Description string	`json:"description" form:"description" gorm:"not null"`
	File		*[]byte	`json:"file" form:"file"`
}

type Submission struct {
	gorm.Model
	AssignmentID int	`json:"assignment_id" form:"assignment_id" gorm:"not null"`
	StudentID    int	`json:"student_id" form:"student_id" gorm:"not null"`
	File         *[]byte	`json:"file" form:"file" gorm:"not null"`
	Comment	  	string	`json:"comment" form:"comment"`
	Grade 	   	int	`json:"grade" form:"grade"`
}

type Otp struct {
	StudentEmail string `json:"student_email" form:"student_email"`
	TeacherEmail string `json:"teacher_email" form:"teacher_email"`
	StudentOTP string `json:"student_otp" form:"student_otp"`
	TeacherOTP string `json:"teacher_otp" form:"teacher_otp"`
}