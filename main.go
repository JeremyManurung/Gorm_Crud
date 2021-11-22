package main

import (
	"github.com/labstack/echo"
	"res-task/config"
	"res-task/route"
)

func main() {
	app := echo.New()
	config.InitDB()
	config.InitMigration()	
	route.NewRoutes(app)
	app.Start(":1234")
}
