package middleware

import (
	"miniproject/constant"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func CreateTeacherToken(teacherID int)(string , error){
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["teacherID"] = teacherID
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	return token.SignedString([]byte(constant.TEACHER_JWT))
}

func CreateStudentToken(studentID int)(string , error){
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["studentID"] = studentID
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	return token.SignedString([]byte(constant.STUDENT_JWT))
}