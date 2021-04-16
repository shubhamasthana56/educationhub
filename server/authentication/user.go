package authentication
import(	
		"jwt-authentication/db")

type User struct {
	Email string 
	Password string 
}

var UserCollection = db.ConnectDB().Database("authentication-test").Collection("user")