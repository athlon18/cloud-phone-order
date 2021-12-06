package service

import (
	"errors"
	"gorm.io/gorm"
	"order/common/util"
	"order/db"
	"order/model"
	"time"
)

func Login(name string, password string) (user model.User, err error) {
	if err = db.DB.Model(model.User{}).Where("name = ? and password = ?", name, password).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return user, errors.New("账号密码不存在，或者密码错误！")
		}
		return user, err
	}
	user.Token = util.GetToken(user.Name, user.Code)
	return user, db.DB.Model(model.User{}).Where("id = ? ", user.ID).Update("token", user.Token).Error

}

func ExternalLoginRegister(name string, password, code string) (err error) {
	user := model.User{}
	if err = db.DB.Model(model.User{}).Where("name = ? and password = ?", name, password).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("账号密码不存在，或者密码错误！")
		}
		return err
	}

	machine := model.UserMachine{
		UserId:  user.ID,
		Machine: code,
	}
	machineDB := db.DB.Model(model.UserMachine{}).Where("machine = ?", code)
	if err = machineDB.First(map[string]interface{}{}).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return db.DB.Model(model.UserMachine{}).Create(&machine).Error
		}
	}

	return machineDB.Update("updated_at", time.Now()).Error
}
