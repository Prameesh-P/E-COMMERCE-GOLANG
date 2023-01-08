package controllers

import (
	"fmt"
	"github.com/Prameesh-P/E-COMMERCE/authentification"
	"github.com/Prameesh-P/E-COMMERCE/initializers"
	"github.com/Prameesh-P/E-COMMERCE/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

var validate = validator.New()

func SignUp(c *gin.Context) {
	var users models.User
	if err := c.ShouldBindJSON(&users); err != nil {
		c.JSON(404, gin.H{
			"msg": err.Error(),
		})
		c.Abort()
		return
	}
	ValidationErr := validate.Struct(users)
	if ValidationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"ValidationError": ValidationErr,
		})
		return
	}
	if err := users.HashPassword(users.Password); err != nil {
		c.JSON(404, gin.H{
			"err": err.Error(),
		})
		c.Abort()
		return
	}
	fmt.Println(users)
	rec := initializers.DB.Create(&users)
	if rec.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": rec.Error.Error(),
		})
	}
	//stmt := `INSERT INTO users(id,first_name,last_name,email,password,phone,block_status,country,city,pincode) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)`
	//initializers.DB.Exec(stmt, users.ID, users.First_Name, users.Last_Name, users.Email, users.Password, users.Phone, users.Block_status, users.Country, users.City)
	c.JSON(200, gin.H{
		"email": users.Email,
		"msg":   "Go to Loginpage",
	})
}

type uLogin struct {
	Email        string
	Password     string
	Block_status bool
}

func Login(c *gin.Context) {
	var ulogin uLogin
	var user models.User
	if err := c.ShouldBindJSON(&ulogin); err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		c.Abort()
		return
	}
	record := initializers.DB.Raw("SELECT * FROM users WHERE email=?", ulogin.Email).Scan(&user)
	if record.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": record.Error.Error(),
		})
		c.Abort()
		return
	}
	if user.Block_status {
		c.JSON(404, gin.H{
			"error": "user has been blocked by admin",
		})
		c.Abort()
		return
	}
	credentialCheck := user.CheckPassword(ulogin.Password)
	if credentialCheck != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg": "Invalid Credentials",
		})
		c.Abort()
		return
	}
	tokenString, err := authentification.GenerateJWT(user.Email)
	fmt.Println(tokenString)
	token := tokenString["access_token"]
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("UserAuth", token, 3600*24*30, "", "", false, true)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{
		"email":    ulogin.Email,
		"password": ulogin.Password,
		"token":    tokenString,
	})
}
func UserHome(c *gin.Context) {
	c.JSON(200, gin.H{
		"success": "Welcome to user home page..!!",
	})
}
