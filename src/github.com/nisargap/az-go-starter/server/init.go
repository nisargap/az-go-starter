// Package server initializes the main backend server with middlewares
// and routes.
package server

import (
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/nisargap/az-go-starter/db"
	"github.com/nisargap/az-go-starter/jwt"
	"github.com/nisargap/az-go-starter/models"
	"github.com/nisargap/az-go-starter/routes"
	"gopkg.in/mgo.v2"
	"os"
	"time"
)

// Given a configuration filename this function
// decodes the json file into the structure defined
// in models.Config in the models package.
func GetConfig(filename string) models.Config {
	file, _ := os.Open(filename)
	decoder := json.NewDecoder(file)
	config := models.Config{}
	err := decoder.Decode(&config)
	fmt.Println(config.DatabaseName)
	if err != nil {
		fmt.Println(err)
	}
	return config
}

// This is the more granular method of configuring Cors
// func ConfigureCors(config models.Config) cors.Config {
// 	corsConfig := cors.Config{
// 		AllowOrigins:     config.AllowedOrigins,
// 		AllowMethods:     []string{"PUT", "PATCH", "OPTIONS", "POST", "GET"},
// 		ExposedHeaders:   "",
// 		ValidateHeaders:  false,
// 		AllowCredentials: true,
// 		MaxAge:           12 * time.Hour,
// 	}
// 	return corsConfig
// }

// Given a filename for the configuration such as config.json
// This function sets up cors using the default config and
// allows all origins
// Sets the database session in the gin config, adds the JWT middleware
// to gin and sets up the routes with the correct API version.
// The routes are split up into two sections those that using the authentication
// middleware and those that don't.
func SetupRouter(filename string) *gin.Engine {
	config := GetConfig(filename)
	r := gin.Default()
	// r.Use(cors.New(ConfigureCors(config)))
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AddAllowHeaders("Authorization")
	r.Use(cors.New(corsConfig))
	url := config.DatabaseUrl
	session, err := mgo.DialWithTimeout(url, time.Duration(5*time.Second))
	if err != nil {
		fmt.Println(err)
	}
	middleware := db.CreateMiddleware(session)
	r.Use(middleware.Connect)
	r.Use(func(c *gin.Context) {
		c.Set("config", config)
		c.Next()
	})
	authMiddleware := jwt.GetJWTMiddleware()
	v1 := r.Group("/api/v1")
	{
		v1.GET("/", routes.MainRoute)
		v1.POST("/login", authMiddleware.LoginHandler)
		v1.POST("/signup", routes.SignUp)
	}
	v1.Use(authMiddleware.MiddlewareFunc())
	{
		v1.GET("/ping", routes.GetPongHandler)
	}
	return r
}
