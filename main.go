package main

import (
	
	// d "github.com/Prameesh-P/E-COMMERCE/database"
	"github.com/Prameesh-P/E-COMMERCE/routes"
	"github.com/gin-gonic/gin"
	"os"
)
func init()  {
	// d.Getenv()
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	router := gin.New()
	router.Use(gin.Logger())

	routes.Authroutes(router)
	routes.Userroutes(router)


	router.Run(":" + port)
}
