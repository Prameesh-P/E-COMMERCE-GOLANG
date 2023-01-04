package controllers

import (
	"context"
	"github.com/Prameesh-P/E-COMMERCE/helpers"
	"github.com/Prameesh-P/E-COMMERCE/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"time"
)

var validate = validator.New()

func Getusers() {

}
func Getuser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("user_id")
		if err := helpers.MatchUserTypeToUid(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
	}
}
func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.DbUser

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		ValidationErr := validate.Struct(user)
		if ValidationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": ValidationErr.Error(),
			})
			return
		}
	}
}
func Login() {

}
