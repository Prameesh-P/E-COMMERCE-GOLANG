package models

import (
	"time"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID            uint      `gorm:"primary key;autoIncrement" json:"id"`
	First_name    string   `json:"first_Name" validate:"required, min=2,max=100"`
	Last_name     string   `json:"last_Name" validate:"required,min=2,max=100"`
	Password      string   `json:"password" validate:"required,min=6"`
	Email         string   `json:"email" validate:"email,required"`
	Phone         string   `json:"phone" validate:"required,min=10"`
	Token         string   `json:"token"`
	User_type     string   `json:"user_type" validate:"required, eq=ADMIN|eq=USER "`
	Refresh_token string   `json:"refresh_token"`
	Created_at    time.Time `json:"created_at"`
	Updated_at    time.Time `json:"updated_at"`
	User_id       string    `json:"user_id"`
}
func (user *User)HashPassword(password string)error  {
	bytes,err := bcrypt.GenerateFromPassword([]byte(password),14)
	if err!=nil{
		return err
	}
	user.Password=string(bytes)
	return nil
}