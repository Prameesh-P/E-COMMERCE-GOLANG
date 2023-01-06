package routes

import (
	c "github.com/Prameesh-P/E-COMMERCE/controllers"
	// "github.com/Prameesh-P/E-COMMERCE/middleware"
	"github.com/gin-gonic/gin"
)

func Userroutes(incomingRoutes *gin.Engine) {

	incomingRoutes.POST("/signup",c.SignUp)

}
