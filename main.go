package main

import (
	"go-app/config"
	"go-app/server"
)


func main() {

	c := config.Config{}

	if err := c.InitDB(); err == nil || config.DB != nil {
		server.InitServer()
	}

	defer c.CloseDB()

}

