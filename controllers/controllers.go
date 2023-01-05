package controllers

import (
	"net/http"
	"github.com/Prameesh-P/E-COMMERCE/database"
	"github.com/Prameesh-P/E-COMMERCE/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

	var validate = validator.New()

func Getusers() {

}

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
			"error":ValidationErr,
		})
		c.Abort()
		return   
	}
	if err := user.HashPassword(user.Password); err!= nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
		c.Abort()
		return
	}
	record:=database.DB.Create(&user)
	if record.Error !=nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"err":record.Error.Error(),
		})
	}


}


func Login() {

}
