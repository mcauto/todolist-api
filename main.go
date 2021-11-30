package main

import (
	"todolist-api/modules/delivery/cmd"
)

// @title todolist API
// @version 21.0.0
// @description todolist API

// @contact.name API Support
// @contact.url http://localhost:5000
// @contact.email nexters@kakao.com

// @host localhost:5000
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	cmd.Execute()
}
