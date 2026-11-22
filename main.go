package main

import (
	"go-social-api/src/config"
	"go-social-api/src/routes"
)

func main() {
	config.InitDB()
	r := routes.SetupRouter()
	r.Run(":9090")
}
