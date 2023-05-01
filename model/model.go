package model

import (
	"github.com/jinzhu/gorm"
)

type Student struct {
	gorm.Model
	Name       string
	Email      string
	Password   string
	Enrollment []Enrollment `gorm:"foreignKey:StudentID"`
	Submission []Submission `gorm:"foreignKey:StudentID"`
}

type Teacher struct {
	gorm.Model
	Name       string
	Email      string
	Password   string
	Classes    []Class `gorm:"foreignKey:TeacherID"`
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
	Assignment  []Assignment `gorm:"foreignKey:ClassID"`
	Material    []Material`gorm:"foreignKey:ClassID"`
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