package routes

import (
	"ormapi/controllers"
	"ormapi/helper"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func Route(e *echo.Echo, uc controllers.UserController, bc controllers.BookController, kc controllers.KeyController, db *gorm.DB) {
	vo := helper.ValidateObj{DB: db}

	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	// e.Use(middleware.LoggerWithConfig(
	// 	middleware.LoggerConfig{
	// 		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	// 	},
	// ))

	users := e.Group("/users", middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		KeyLookup: "header:api-key",
		Validator: vo.ValidateKey,
	}))
	books := e.Group("/books", middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		KeyLookup: "header:api-key",
		Validator: vo.ValidateKey,
	}))

	e.GET("/mykeys", kc.AskKey())
	e.POST("/login", uc.Login())

	users.POST("", uc.Register)
	users.GET("", uc.GetUser())
	// e.GET("/users/:user_id/books")

	books.GET("/:bookId", bc.GetBookByID())
	books.GET("", bc.GetBook())
	books.POST("", bc.AddBook)

}
