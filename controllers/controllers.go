package controllers

import (
	"github.com/Prameesh-P/E-COMMERCE/database"

	"net/http"
	"github.com/Prameesh-P/E-COMMERCE/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func SignUp(c *gin.Context){ 
	var user models.User
	if err:=c.ShouldBindJSON(&user);err!=nil{
		c.JSON(404,gin.H{
			"msg":err.Error(),
		})
		c.Abort()
		return
	}
	ValidationErr:=validate.Struct(user)
	if ValidationErr!=nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"ValidationError":ValidationErr,
		})
		c.Abort()
		return   
	}
	if err := user.HashPassword(user.Password); err!= nil {
		c.JSON(404,gin.H{
			"err":err.Error(),
		})
		c.Abort()
		return
	}
	record:=.DB.Create(&user)
	c.JSON(200,gin.H{
		"email":user.Email,
		"msg":"Go to Loginpage",
	})
}

func Login() {

}
