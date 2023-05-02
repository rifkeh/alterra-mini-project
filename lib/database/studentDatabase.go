package database

import (
	"miniproject/config"
	"miniproject/middleware"
	"miniproject/model"
)

func LoginStudent(student *model.Student) (interface{}, error) {
	var err error
	if err = config.DB.Where("email = ? AND password = ?", student.Email, student.Password).First(&student).Error; err != nil {
		return nil, err
	}
	studentResp := map[string]interface{}{
		"Name":  student.Name,
		"Email": student.Email,
		"Token" : "",
	}
	studentResp["Token"], err = middleware.CreateStudentToken(int(student.ID))
	if err != nil {
		return nil, err
	}
	return studentResp, nil
}