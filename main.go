package main

import (
	bhandle "ormapi/app/features/book/handler"
	brepo "ormapi/app/features/book/repository"
	blogic "ormapi/app/features/book/usecase"
	uhandle "ormapi/app/features/user/handler"
	urepo "ormapi/app/features/user/repository"
	ulogic "ormapi/app/features/user/usecase"
	"ormapi/app/routes"
	"ormapi/config"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	cfg := config.InitSQL()
	cfg.AutoMigrate(urepo.User{})
	cfg.AutoMigrate(brepo.Book{})

	mdl := urepo.New(cfg)
	srv := ulogic.New(mdl)
	ctl := uhandle.New(srv)

	bookMdl := brepo.New(cfg)
	bookSrv := blogic.New(bookMdl)
	bookCtl := bhandle.New(bookSrv)

	// ROUTING
	routes.Route(e, ctl, bookCtl)

	e.Start(":8000")
}
