package authentication
import( "context"
		"fmt"
		"log"

		"go.mongodb.org/mongo-driver/bson"
		"go.mongodb.org/mongo-driver/mongo/options")


func IsUserExist(userData User) bool {
cur,err := UserCollection.Find(context.TODO(), bson.D{{}}, options.Find())
//var users []*user
if err != nil {
	log.Fatal(err)
}
for cur.Next(context.TODO()) {
	var result user
	err = cur.Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	if userData.Email == result.Email {
		return true
	}
	// fmt.Println(result.Email, "result")
	// if err != nil {  }
	// // do something with result....
	// users = append(users, &result)
 }
 return false
}