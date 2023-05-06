package database

import (
	"miniproject/config"
	"miniproject/middleware"
	"miniproject/model"
)

func LoginTeacher(teacher *model.Teacher) (map[string]string, error) {
	var err error
	if err = config.DB.Where("email = ? AND password = ?", teacher.Email, teacher.Password).First(&teacher).Error; err != nil {
		return nil, err
	}
	teacherResp := map[string]string{
		"Name":  teacher.Name,
		"Email": teacher.Email,
		"Token" : "",
	}
	teacherResp["Token"], err = middleware.CreateTeacherToken(int(teacher.ID))
	if err != nil {
		return nil, err
	}
	// teacher.LastLogin = config.DB.NowFunc()
	if err := config.DB.Save(&teacher).Error; err != nil {
		return nil, err
	}
	return teacherResp, nil
}