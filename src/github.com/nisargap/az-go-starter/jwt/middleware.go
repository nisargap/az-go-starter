package jwt

import (
	"fmt"
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"github.com/nisargap/az-go-starter/models"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

func GetJWTMiddleware() (j *jwt.GinJWTMiddleware) {
	const notFound = "NOTFOUND"
	const bcryptCost = 10 // cost for bcrypto algo higher == more secure
	// the jwt middleware
	authMiddleware := &jwt.GinJWTMiddleware{
		Realm:      "test zone",
		Key:        []byte("secret key"),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		Authenticator: func(email string, password string, c *gin.Context) (string, bool) {
			// Get the salt from the database for the particular email
			// Use bcrypt and hash the password with the retrieved salt
			// Check if the resultant hash matches with the one in the database
			// If the hash does match return the user ObjectId and true
			// Else return the NOTFOUND and false
			config := c.MustGet("config").(models.Config)
			// Old Code: users := db.GetCollection(config.DatabaseName, "users", c)
			session := c.MustGet("db").(*mgo.Session)
			users := session.DB(config.DatabaseName).C("users")
			var userFound models.User
			err := users.Find(bson.M{"email": email}).One(&userFound)
			if err != mgo.ErrNotFound {
				salt := userFound.Salt
				passwordWithSalt := password + salt
				if err != nil {
					fmt.Println(err)
				}
				authErr := bcrypt.CompareHashAndPassword([]byte(userFound.Password), []byte(passwordWithSalt))
				if authErr == nil {
					return userFound.Id.Hex(), true
				}
			}
			return notFound, false
		},
		Authorizator: func(userId string, c *gin.Context) bool {
			// Check if the ObjectId is in the database if it is return true
			// Else return False
			config := c.MustGet("config").(models.Config)
			// Old code: users := db.GetCollection(config.DatabaseName, "users", c)
			session := c.MustGet("db").(*mgo.Session)
			users := session.DB(config.DatabaseName).C("users")
			if userId == "" {
				return false
			}
			num, err := users.Find(bson.M{"_id": bson.ObjectIdHex(userId)}).Count()
			if err != nil {
				fmt.Println(err)
			}
			if num == 1 {
				return true
			}
			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		TokenLookup: "header:Authorization",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	}
	return authMiddleware

}
