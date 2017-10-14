package jwt

import (
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"github.com/nisargap/az-go-starter/db"
	"github.com/nisargap/az-go-starter/models"
)

func GetCurrentIdentity(c *gin.Context) models.User {
	// get the current claims
	claims := jwt.ExtractClaims(c)
	id := claims["id"].(string)
	// db.GetUser(id) -> User
	user := db.GetUser(id, c)
	return user
}
