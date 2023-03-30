package helper

import (
	"ormapi/entities"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ValidateObj struct {
	DB *gorm.DB
}

func (vo *ValidateObj) ValidateKey(key string, c echo.Context) (bool, error) {
	var keyResult entities.Keys
	err := vo.DB.Select("key").Where("`key` = ?", key).First(&keyResult).Error
	if err != nil {
		return false, nil
	}

	return key == keyResult.Key, nil
}
