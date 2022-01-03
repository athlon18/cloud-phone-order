package service

import (
	"errors"
	uuid "github.com/satori/go.uuid"
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

func ExternalLoginRegister(name string, password, code string) (userMachine model.UserMachine, err error) {
	user := model.User{}
	if err = db.DB.Model(model.User{}).Where("name = ? and password = ?", name, password).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return userMachine, errors.New("账号密码不存在，或者密码错误！")
		}
		return
	}
	machineData := new(model.Machine)
	// 机器码发现检测
	if err = db.DB.Model(model.Machine{}).Where("machine_code = ?", code).First(machineData).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return userMachine, errors.New("机器码未注册！请先注册！")
		}
		return
	}

	userMachine = model.UserMachine{
		UserId:    user.ID,
		Machine:   uuid.NewV4().String(),
		MachineId: machineData.ID,
	}
	machineDB := db.DB.Model(model.UserMachine{}).Where("machine_id = ? and user_id = ?", machineData.ID, user.ID)
	if err = machineDB.First(map[string]interface{}{}).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return userMachine, db.DB.Model(model.UserMachine{}).Create(&userMachine).Error
		}
	}

	return userMachine, machineDB.Update("updated_at", time.Now()).Error
}

func GetMachineInfo(code string) (machineData model.UserMachine, err error) {
	// 机器码发现检测
	if err = db.DB.Model(model.UserMachine{}).Where("machine = ?", code).First(machineData).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return machineData, errors.New("用户机器码过期！请重新登录！")
		}
		return
	}
	return machineData, err
}
