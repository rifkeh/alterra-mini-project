package database

import (
	"fmt"
	"math/rand"
	"miniproject/config"
	"miniproject/middleware"
	"miniproject/model"
	"time"

	sendinblue "github.com/CyCoreSystems/sendinblue"
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
	// student.LastLogin = config.DB.NowFunc()
	if err := config.DB.Save(&student).Error; err != nil {
		return nil, err
	}
	return studentResp, nil
}


func SendEmail(toName, toEmail , otp string) error {
	sender := sendinblue.Address{
		Name:  "Miniproject",
		Email: "test@example.com",
	}
	recipient := sendinblue.Address{
		Name:  toName,
		Email: toEmail,
	}
    message := sendinblue.Message{
		Sender: &sender,
		To:     []*sendinblue.Address{&recipient},
		Subject: "Account Creation",
		TextContent:    fmt.Sprintf("OTP: %s", otp),
	}
	return message.Send("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")
}

func GenerateToken() string{
	rand.Seed(time.Now().UnixNano())
	chars := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	otpBytes := make([]byte, 6)
	for i := range otpBytes {
		otpBytes[i] = chars[rand.Intn(len(chars))]
	}
	return string(otpBytes)
}