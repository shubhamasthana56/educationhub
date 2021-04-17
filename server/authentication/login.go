package authentication

import (
	"context"
	"fmt"
	"jwt-authentication/model"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type LoginResponse struct {
	User          bool `json: "user"`
	Authenticated bool `json: "authenticated"`
}

func LoginController(c *gin.Context) {
	var userResponse model.UserModel
	userParam := User{}
	loginResponse := LoginResponse{}
	userParam.Email = c.Request.FormValue("email")
	userParam.Password = c.Request.FormValue("password")
	cur, err := UserCollection.Find(context.TODO(), bson.M{"email": userParam.Email}, options.Find())
	if err != nil {

	}
	for cur.Next(context.TODO()) {

		err = cur.Decode(&userResponse)
		if err != nil {
			log.Fatal(err)
		}
	}
	loginResponse = verifyUserHash(userResponse, userParam)
	fmt.Println(loginResponse)

}

func verifyUserHash(response model.UserModel, user User) LoginResponse {
	if len(strings.TrimSpace(response.Email)) == 0 {
		return LoginResponse{User: false, Authenticated: false}
	} else {
		// Comparing the password with the hash
		err := bcrypt.CompareHashAndPassword(response.Password, []byte(user.Password))
		if err == nil {
			return LoginResponse{User: true, Authenticated: true}
		}
	}
	return LoginResponse{User: true, Authenticated: false}
}
