package main

import (
	"fmt"
	"ormapi/book"
	"ormapi/config"
	"ormapi/helper"
	"ormapi/routes"
	"ormapi/user"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	cfg := config.InitSQL()
	cfg.AutoMigrate(user.User{})
	cfg.AutoMigrate(book.Book{})

	mdl := user.UserModel{}
	mdl.SetDB(cfg)
	ctl := user.UserController{}
	ctl.SetModel(mdl)

	bookMdl := book.BookModel{}
	bookMdl.SetDB(cfg)
	bookCtl := book.BookController{}
	bookCtl.SetModel(bookMdl)

	token := helper.GenerateJWT(1, "manager")
	fmt.Println("jwt yang dibuat", token)

	// ROUTING
	routes.Route(e, ctl, bookCtl)

	e.Start(":8000")
}
