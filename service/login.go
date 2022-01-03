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
	tx := db.DB.Begin()
	oldUserMachine := model.UserMachine{}
	machineDB := tx.Model(model.UserMachine{}).Where("machine_id = ?", machineData.ID)
	if err = machineDB.First(&oldUserMachine).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			if err = tx.Model(model.UserMachine{}).Create(&userMachine).Error; err != nil {
				tx.Rollback()
				return
			}
			return userMachine, tx.Commit().Error
		}
		tx.Rollback()
		return
	}
	// 检测machine id 归属
	if oldUserMachine.UserId != user.ID {
		tx.Rollback()
		return userMachine, errors.New("此机器已被其他用户绑定！")
	}

	//处理 之前冗余数据
	if err = tx.Model(model.Order{}).
		Where("machine = ?", oldUserMachine.Machine).
		Update("machine", userMachine.Machine).Error; err != nil {
		tx.Rollback()
		return
	}

	// 更新时间戳
	if err = machineDB.Update("updated_at", time.Now()).Error; err != nil {
		tx.Rollback()
		return
	}

	return userMachine, tx.Commit().Error
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
