package routes

import (
	"github.com/gin-gonic/gin"
)

// Route handler
func GetPongHandler(c *gin.Context) {
	c.String(200, "pong")
	return
}
