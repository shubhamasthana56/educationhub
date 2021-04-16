package authentication

import (
	"context"
	"strings"
	"log"
	"jwt-authentication/model"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	
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

func userAdd(c *gin.Context, userData User) {
	    if IsUserExist(userData) {
			c.JSON(400, gin.H{
				"message": "User Exist with" + userData.Email,
		})
		} else {
			// Hashing the password with the default cost of 10
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userData.Password), bcrypt.DefaultCost)
			userSignupData := model.UserModel{}
			userSignupData.Email = userData.Email
			userSignupData.Password = hashedPassword
			if err != nil {
				log.Fatal(err)
			}
			_, err = UserCollection.InsertOne(context.TODO(), userSignupData)
			if err != nil {
				log.Fatal(err)
			}
			c.JSON(200, gin.H{
				"message": "User added successfully",
			})
		}
		
}
