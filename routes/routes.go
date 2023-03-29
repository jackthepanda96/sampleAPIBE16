package routes

import (
	"ormapi/book"
	"ormapi/user"

	"github.com/labstack/echo/v4"
)

func Route(e *echo.Echo, uc user.UserController, bc book.BookController) {
	e.POST("/users", uc.Register)
	e.POST("/login", uc.Login())
	e.GET("/users", uc.GetUser())
	// e.GET("/users/:user_id/books")

	e.GET("/books/:bookId", bc.GetBookByID())
	e.GET("/books", bc.GetBook())
	e.POST("/books", bc.AddBook)
}
