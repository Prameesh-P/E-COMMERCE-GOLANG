package models

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID           uint   `json:"id" gorm:"primaryKey;unique"		 `
	First_Name   string `json:"first_name"   validate:"required,min=2,max=50"  `
	Last_Name    string `json:"last_name"       validate:"required,min=1,max=50"  `
	Email        string `json:"email"     validate:"email,required"`
	Password     string `json:"password"   validate:"required"`
	Phone        string `json:"phone"    validate:"required"`
	Block_status bool   `json:"block_status "    `
	Country      string `json:"country "   `
	City         string `json:"city "   `
	Pincode      uint   `json:"pincode "   `
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
type Admin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (admin *Admin) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	admin.Password = string(bytes)
	return nil
}
func (admin *Admin) CheckPassword(incomingPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(incomingPassword))
	if err != nil {
		return err
	}
	return nil
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}
func (user *User) CheckPassword(ProvidedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(ProvidedPassword))
	if err != nil {
		return err
	}
	return nil
}
