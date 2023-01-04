package routes

import (
	"github.com/Prameesh-P/E-COMMERCE/controllers"
	"github.com/gin-gonic/gin"
)

func Authroutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/signup", controllers.SignUp())
	incomingRoutes.POST("users/login", controllers.Login())
}
