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

	users := e.Group("/users")
	books := e.Group("/books")

	e.POST("/login", uc.Login())

	users.POST("/users", uc.Register)
	users.GET("/users", uc.GetUser())
	// e.GET("/users/:user_id/books")

	books.GET("/books/:bookId", bc.GetBookByID())
	books.GET("/books", bc.GetBook())
	books.POST("/books", bc.AddBook)
}
