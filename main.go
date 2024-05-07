package main

import (
	"github.com/labstack/echo/v4"
	"github.com/satryanararya/go-chefbot/config"
	"github.com/satryanararya/go-chefbot/drivers/database"
)

func main() {
	config.LoadEnv()
	config.InitConfigDB()
	DB := database.ConnectDB(config.InitConfigDB())

	e := echo.New()

	e.Logger.Fatal(e.Start(":1323"))
}