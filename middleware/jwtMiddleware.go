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

func ExtractStudentIdToken(token string)(float64){
	claims := jwt.MapClaims{}
	tempToken , _ := jwt.ParseWithClaims(token,claims,func(tempToken *jwt.Token)(interface{},error){
		return []byte(constant.STUDENT_JWT),nil
	},
	)
	return tempToken.Claims.(jwt.MapClaims)["studentID"].(float64)
}