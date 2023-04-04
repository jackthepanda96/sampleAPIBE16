package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	// cfg := config.InitSQL()
	// cfg.AutoMigrate(user.User{})
	// cfg.AutoMigrate(book.Book{})

	// mdl := user.UserModel{}
	// mdl.SetDB(cfg)
	// ctl := user.UserController{}
	// ctl.SetModel(mdl)

	// bookMdl := book.BookModel{}
	// bookMdl.SetDB(cfg)
	// bookCtl := book.BookController{}
	// bookCtl.SetModel(bookMdl)

	// // ROUTING
	// routes.Route(e, ctl, bookCtl)

	e.GET("/hello", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello world")
	})

	e.Start(":8000")
}
