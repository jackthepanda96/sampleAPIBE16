package main

import (
	"ormapi/config"
	"ormapi/controllers"
	"ormapi/entities"
	"ormapi/models"
	"ormapi/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	cfg := config.InitSQL()
	cfg.AutoMigrate(entities.User{})
	cfg.AutoMigrate(entities.Book{})
	cfg.AutoMigrate(entities.Keys{})

	mdl := models.UserModel{}
	mdl.SetDB(cfg)
	ctl := controllers.UserController{}
	ctl.SetModel(mdl)

	bookMdl := models.BookModel{}
	bookMdl.SetDB(cfg)
	bookCtl := controllers.BookController{}
	bookCtl.SetModel(bookMdl)

	keyMdl := models.KeyModel{}
	keyMdl.SetModel(cfg)
	keyCtl := controllers.KeyController{}
	keyCtl.SetModel(keyMdl)
	ctl.Km = keyMdl

	// ROUTING
	routes.Route(e, ctl, bookCtl, keyCtl, cfg)

	e.Start(":8000")
}
