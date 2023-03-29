package user

import (
	"ormapi/book"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nama     string `json:"nama"`
	HP       string `json:"hp" gorm:"type:varchar(13);primaryKey"`
	Password string `json:"password"`
	Book     []book.Book
}
