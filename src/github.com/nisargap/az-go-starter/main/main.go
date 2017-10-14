// Package main sets up the router using config.json
// Gets the configuration from config.json and
// uses that to run the server usig the SSL info locations.
package main

import (
	"github.com/nisargap/az-go-starter/server"
)

var DB = make(map[string]string)

func main() {
	router := server.SetupRouter("config.json")
	config := server.GetConfig("config.json")
	router.Run(":" + config.Port)
}
