package controllers

import (
	"github.com/Prameesh-P/E-COMMERCE/authentification"
	i "github.com/Prameesh-P/E-COMMERCE/initializers"
	"github.com/Prameesh-P/E-COMMERCE/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AdminLogins struct {
	Email    string
	Password string
}

var UserDB = map[string]string{
	"email":    "prameepramee0@gmail.com",
	"password": "pramee123",
}

func AdminSignup(c *gin.Context) {
	var admin models.Admin
	var count uint
	if err := c.ShouldBindJSON(&admin); err != nil {
		c.JSON(404, gin.H{
			"err": err.Error(),
		})
		c.Abort()
		return
	}
	i.DB.Raw("SELECT count(*) FROM admins WHERE email=?", admin.Email).Scan(&count)
	if count > 0 {
		c.JSON(404, gin.H{
			"err": "false",
			"msg": "Admin with same Email already exists",
		})
		c.Abort()
		return
	}
	if err := admin.HashPassword(admin.Password); err != nil {
		c.JSON(500, gin.H{
			"err": err.Error(),
		})
	}
	record := i.DB.Create(&admin)
	if record.Error != nil {
		c.JSON(400, gin.H{
			"error": record.Error.Error(),
		})
	}
	c.JSON(200, gin.H{
		"status": "OK",
		"msg":    "new admin created",
	})
}
func AdminLogin(c *gin.Context) {
	var adminLogin AdminLogins
	var admin models.Admin
	if err := c.ShouldBindJSON(&adminLogin); err != nil {
		c.JSON(404, gin.H{
			"err": err.Error(),
		})
		c.Abort()
		return
	}
	var record = i.DB.Raw("SELECT * FROM admins WHERE email=?", adminLogin.Email).Scan(&admin)
	if record.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": record.Error.Error(),
		})
		c.Abort()
		return
	}
	credentialCheck := admin.CheckPassword(adminLogin.Password)
	if credentialCheck != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid Credentials",
		})
		c.Abort()
		return
	}
	tokenString, err := authentification.GenerateJWT(adminLogin.Email)
	token := tokenString["access_token"]
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("AdminJWT", token, 3600*24*30, "", "", false, true)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{
		"status":      "true",
		"msg":         "OK",
		"tokenstring": tokenString,
	})
}
func AdminHome(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "welcome to admin Home page",
	})
}
