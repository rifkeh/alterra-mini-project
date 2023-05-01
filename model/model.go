package model

import (
	"github.com/jinzhu/gorm"
)

type Student struct {
	gorm.Model
	Name       string
	Email      string
	Password   string
	Enrollment []Enrollment
}

type Teacher struct {
	gorm.Model
	Name       string
	Email      string
	Password   string
	Enrollment []Enrollment
	Classes    []Class
}

type Enrollment struct {
	gorm.Model
	StudentID uint
	ClassID   uint
}

type Class struct {
	gorm.Model
	TeacherID   uint
	Name        string
	Description string
	Assignment  []Assignment
	Material    []Material
}

type Assignment struct {
	gorm.Model
	ClassID     uint
	Description string
	Deadline    string
}

type Material struct {
	gorm.Model
	ClassID     uint
	Description string
}

type Submission struct {
	gorm.Model
	AssignmentID uint
	StudentID    uint
	Link         string
}