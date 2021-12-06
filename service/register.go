package service

import (
	"errors"
	"gorm.io/gorm"
	"order/common/util"
	"order/db"
	"order/model"
)

func Register(name string, password string) (model.User, error) {
	code := util.GetRandomString(6)
	user := model.User{
		Name:     name,
		Password: util.EncryptSha256(password),
		Code:     code,
		Token:    util.GetToken(name, code),
	}
	if err := db.DB.Model(model.User{}).Where("name = ?", name).First(map[string]interface{}{}).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			if err = db.DB.Model(model.User{}).Create(&user).Error; err != nil {
				return user, errors.New("注册账号失败！")
			}
			return user, nil
		}
		return user, err
	}
	return user, errors.New("账号已存在！")
}
