package models

import (
	"errors"
	"log"
	"ormapi/entities"

	"gorm.io/gorm"
)

type UserModel struct {
	db *gorm.DB
}

func (um *UserModel) SetDB(db *gorm.DB) {
	um.db = db
}

func (um *UserModel) Insert(newUser entities.User) (entities.User, error) {
	if err := um.db.Create(&newUser).Error; err != nil {
		log.Println("Terjadi error saat create user", err.Error())
		return entities.User{}, err
	}

	return newUser, nil
}

func (um *UserModel) Login(hp, password string) (entities.User, error) {
	res := entities.User{}
	// Query login -> select * from users where hp = ? and password = ?
	if err := um.db.Where("hp = ? AND password = ?", hp, password).Find(&res).Error; err != nil {
		log.Println("Terjadi error saat create user", err.Error())
		return entities.User{}, err
	}

	if res.HP == "" {
		log.Println("Data tidak ditemukan")
		return entities.User{}, errors.New("data tidak ditemukan")
	}

	return res, nil
}

func (um *UserModel) GetAllUser() ([]entities.User, error) {
	res := []entities.User{}

	if err := um.db.Select("hp, nama, id").Find(&res).Error; err != nil {
		log.Println("Terjadi error saat select user ", err.Error())
		return nil, err
	}

	return res, nil
}
