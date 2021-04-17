package main

import (
	"jwt-authentication/authentication"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/signup", authentication.UserAddController)
	router.POST("/login", authentication.LoginController)
	router.Run(":8080")

}
