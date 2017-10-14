package db

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
)

func GetSession(c *gin.Context) *mgo.Session {
	session := c.MustGet("db").(*mgo.Session)
	session.Refresh()
	return session.Copy()
}
