package service

import (
	"errors"
	"order/db"
	"order/model"
)

func GetTokenInfo(token string) (model.User, error) {
	user := model.User{}
	if err := db.DB.Model(model.User{}).Where("token = ?", token).First(&user).Error; err != nil {
		return user, errors.New("令牌已失效，请重新登录！")
	}
	return user, nil
}

func GetCodeInfo(code string) (model.User, error) {
	user := model.User{}
	if err := db.DB.Model(model.User{}).Where("code = ?", code).First(&user).Error; err != nil {
		return user, errors.New("特征码不存在，请重试！")
	}
	return user, nil
}

func GetIDInfo(id uint) (user model.User, err error) {
	if err = db.DB.Model(model.User{}).Where("id = ?", id).First(&user).Error; err != nil {
		return user, errors.New("用户不存在，请重试！")
	}
	return user, nil
}

func UpdateUserByCode(user model.User) error {
	if err := db.DB.Model(model.User{}).Where("code = ?", user.Code).Updates(&user).Error; err != nil {
		return err
	}
	return nil
}

func UpdateUserTokenByID(id int) error {
	return db.DB.Model(model.User{}).Where("id = ? ", id).Update("token", "").Error
}
