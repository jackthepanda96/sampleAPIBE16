package handler

import (
	"net/http"
	"ormapi/app/features/user"
	"ormapi/helper"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	service user.UseCase
}

func New(us user.UseCase) user.Handler {
	return &UserController{
		service: us,
	}
}

func (uc *UserController) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := RegisterInput{}
		if err := c.Bind(&input); err != nil {
			c.Logger().Error("terjadi kesalahan bind", err.Error())
			return c.JSON(helper.ReponsFormat(http.StatusBadRequest, "terdapat kesalahan input dari user", nil))
		}

		err := uc.service.Register(user.Core{HP: input.HP, Nama: input.Nama, Password: input.Password})

		if err != nil {
			return c.JSON(helper.ReponsFormat(http.StatusInternalServerError, err.Error(), nil))
		}

		return c.JSON(helper.ReponsFormat(http.StatusCreated, "sukses menambahkan data", nil))
	}
}

func (uc *UserController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input LoginInput
		if err := c.Bind(&input); err != nil {
			c.Logger().Error("terjadi kesalahan bind", err.Error())
			return c.JSON(helper.ReponsFormat(http.StatusBadRequest, "terdapat kesalahan input dari user", nil))
		}

		res, err := uc.service.Login(input.Hp, input.Password)
		if err != nil {
			return c.JSON(helper.ReponsFormat(http.StatusInternalServerError, err.Error(), nil))
		}

		var result = new(LoginResponse)
		token := helper.GenerateJWT(res.HP)
		result.Nama = res.Nama
		result.HP = res.HP
		result.Token = token

		return c.JSON(helper.ReponsFormat(http.StatusOK, "sukses login, gunakan token ini pada akses api selanjutnya : ", result))
	}
}

// func (uc *UserController) GetUser() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		res, err := uc.model.GetAllUser()

// 		if err != nil {
// 			c.Logger().Error("user model error ", err.Error())
// 			return c.JSON(helper.ReponsFormat(http.StatusInternalServerError, "terdapat kesalahan pada server", nil))
// 		}

// 		return c.JSON(helper.ReponsFormat(http.StatusOK, "sukses menampilkan data", res))
// 	}
// }
