package db

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nisargap/az-go-starter/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func GetUser(hexId string, c *gin.Context) models.User {
	config := c.MustGet("config").(models.Config)
	// Old code: users := GetCollection(config.DatabaseName, "users", c)
	usersSession := c.MustGet("db").(*mgo.Session)
	users := usersSession.DB(config.DatabaseName).C("users")
	var respUser models.User
	fields := bson.M{
		"username":    1,
		"privilege":   1,
		"profile_img": 1,
	}
	err := users.Find(bson.M{"_id": bson.ObjectIdHex(hexId)}).Select(fields).One(&respUser)
	if err != nil {
		fmt.Println(err)
	}
	return respUser
}
