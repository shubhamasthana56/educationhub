package authentication

import (
	"context"
	"fmt"
	"strings"
	"log"

	"github.com/gin-gonic/gin"
	
)

func UserAddController(c *gin.Context) {
	user := User{}
	user.Email = c.Request.FormValue("email")
	user.Password = c.Request.FormValue("password")
	if len(strings.TrimSpace(user.Email)) == 0 {
		c.JSON(400, gin.H{
			"message": "Bad request",
		})
	} else {
		userAdd(c, user)
	}
	

}

func userAdd(c *gin.Context, userData user) {
	    if IsUserExist(userData) {
			c.JSON(400, gin.H{
				"message": "User Exist with" + userData.Email,
		})
		} else {
			_, err := UserCollection.InsertOne(context.TODO(), userData)
			if err != nil {
				log.Fatal(err)
			}
			c.JSON(200, gin.H{
				"message": "User added successfully",
			})
		}
		
}
