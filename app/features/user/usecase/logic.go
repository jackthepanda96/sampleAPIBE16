package usecase

import (
	"errors"
	"ormapi/app/features/user"

	"github.com/labstack/gommon/log"
)

type userLogic struct {
	m user.Repository
}

func New(r user.Repository) user.UseCase {
	return &userLogic{
		m: r,
	}
}

func (ul *userLogic) Login(hp string, password string) (user.Core, error) {
	result, err := ul.m.Login(hp, password)
	if err != nil {
		return user.Core{}, errors.New("terdapat permasalahan pada server")
	}

	return result, nil
}

func (ul *userLogic) Register(newUser user.Core) error {
	_, err := ul.m.Insert(newUser)
	if err != nil {
		log.Error("register logic error:", err.Error())
		return errors.New("terjadi kesalahn pada server")
	}

	return nil
}
