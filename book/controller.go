package book

import (
	"net/http"
	"ormapi/helper"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type BookController struct {
	model BookModel
}

func (bc *BookController) SetModel(m BookModel) {
	bc.model = m
}

func (bc *BookController) AddBook(c echo.Context) error {
	userID, _ := helper.DecodeJWT(c.Get("user").(*jwt.Token))

	input := Book{}
	if err := c.Bind(&input); err != nil {
		c.Logger().Error("terjadi kesalahan bind", err.Error())
		return c.JSON(helper.ReponsFormat(http.StatusBadRequest, "terdapat kesalahan input dari Book", nil))
	}
	input.UserID = userID
	res, err := bc.model.Insert(input)

	if err != nil {
		c.Logger().Error("terjadi kesalahan", err.Error())
		return c.JSON(helper.ReponsFormat(http.StatusInternalServerError, "terdapat kesalahan pada server", nil))
	}

	return c.JSON(helper.ReponsFormat(http.StatusCreated, "sukses menambahkan data", res))
}

func (bc *BookController) GetBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := c.QueryParam("user_id")
		cnv := 0
		if userID != "" {
			cnvID, err := strconv.Atoi(userID)
			cnv = cnvID
			if err != nil {
				c.Logger().Error("Input error ", err.Error())
				return c.JSON(helper.ReponsFormat(http.StatusBadRequest, "terdapat kesalahan pada input ID", nil))
			}
		}

		res, err := bc.model.GetAllBook(uint(cnv))

		if err != nil {
			c.Logger().Error("Book model error ", err.Error())
			return c.JSON(helper.ReponsFormat(http.StatusInternalServerError, "terdapat kesalahan pada server", nil))
		}

		return c.JSON(helper.ReponsFormat(http.StatusOK, "sukses menampilkan data", res))
	}
}

func (bc *BookController) GetBookByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		inputParameter := c.Param("bookId")
		cnv, err := strconv.Atoi(inputParameter)
		if err != nil {
			c.Logger().Error("Input error ", err.Error())
			return c.JSON(helper.ReponsFormat(http.StatusBadRequest, "terdapat kesalahan pada input ID", nil))
		}
		res, err := bc.model.GetBookByID(uint(cnv))

		if err != nil {
			c.Logger().Error("Book model error ", err.Error())
			return c.JSON(helper.ReponsFormat(http.StatusInternalServerError, "terdapat kesalahan pada server", nil))
		}

		return c.JSON(helper.ReponsFormat(http.StatusOK, "sukses menampilkan data", res))
	}
}
