// Package routes contains all of the routes for the API
package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func MainRoute(c *gin.Context) {
	c.String(http.StatusOK, "Works correctly!")
	return
}
