package service

import (
	"errors"
	"gorm.io/gorm"
	"order/common/util"
	"order/db"
	"order/model"
)

func Login(name string, password string) (model.User, error) {
	user := model.User{}
	if err := db.DB.Model(model.User{}).Where("name = ? and password = ?", name, password).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return user, errors.New("账号密码不存在，或者密码错误！")
		}
		return user, err
	}
	user.Token = util.GetToken(user.Name, user.Code)
	db.DB.Model(model.User{}).Where("id = ? ", user.ID).Update("token", user.Token)
	return user, nil
}
