package main

import (
	"github.com/labstack/echo/v4"
	"github.com/satryanararya/go-chefbot/config"
	"github.com/satryanararya/go-chefbot/drivers/database"
	"github.com/satryanararya/go-chefbot/routes"
	"github.com/satryanararya/go-chefbot/utils/validation"
	"gorm.io/gorm"
)

var db *gorm.DB
var v *validation.Validator

func init() {
	config.LoadEnv()
	config.InitConfigDB()
	db = database.ConnectDB(config.InitConfigDB())
	v = validation.NewValidator()
}

func main() {
	e := echo.New()

	routes.InitRoute(e, db, v)

	e.Logger.Fatal(e.Start(":8080"))
}