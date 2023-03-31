package routes

import (
	"ormapi/book"
	"ormapi/user"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Route(e *echo.Echo, uc user.UserController, bc book.BookController) {
	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.POST("/login", uc.Login())
	e.POST("/users", uc.Register)

	e.GET("/books", bc.GetBook())

	e.GET("/users", uc.GetUser(), middleware.JWT([]byte("S3cr3t!!")))
	// // e.GET("/users/:user_id/books")
	e.GET("/books/:bookId", bc.GetBookByID(), middleware.JWT([]byte("S3cr3t!!")))
	e.POST("/books", bc.AddBook, middleware.JWT([]byte("S3cr3t!!")))
}
