package routes

import (
	"ormapi/app/features/book"
	"ormapi/app/features/user"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Route(e *echo.Echo, uc user.Handler, bc book.Handler) {
	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.POST("/login", uc.Login())
	e.POST("/users", uc.Register())

	e.GET("/books", bc.GetAllBook())

	// e.GET("/users", uc.GetUser(), middleware.JWT([]byte("S3cr3t!!")))
	// // e.GET("/users/:user_id/books")

	e.POST("/books", bc.AddBook(), middleware.JWT([]byte("S3cr3t!!")))
}
