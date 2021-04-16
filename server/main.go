package main

import (
	"github.com/gin-gonic/gin"
	"jwt-authentication/authentication"
)

func main() {
	router := gin.Default()
	router.POST("/signup", authentication.UserAddController)
	router.Run(":8080")

}
