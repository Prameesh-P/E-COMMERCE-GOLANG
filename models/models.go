package models

import (
	"time"
	"golang.org/x/crypto/bcrypt"
)


type User struct {
	ID             uint   `json:"id" gorm:"primaryKey;unique"  `
	First_Name     string `json:"first_name"  gorm:"not null" validate:"required,min=2,max=50"  `
	Last_Name      string `json:"last_name"    gorm:"not null"    validate:"required,min=1,max=50"  `
	Email          string `json:"email"   gorm:"not null;unique"  validate:"email,required"`
	Password       string `json:"password" gorm:"not null"  validate:"required"`
	Phone          string `json:"phone"   gorm:"not null;unique" validate:"required"`
	Block_status   bool   `json:"block_status " gorm:"not null"   `
	Country        string `json:"country "   `
	City           string `json:"city "   `
	Pincode        uint   `json:"pincode "   `
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
func (user *User)HashPassword(password string)error  {
	bytes,err := bcrypt.GenerateFromPassword([]byte(password),14)
	if err!=nil{
		return err
	}
	user.Password=string(bytes)
	return nil
}