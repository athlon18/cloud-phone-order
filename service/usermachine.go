package service

import (
	"order/db"
	"order/model"
)

func UserMachineList(where model.UserMachine, userId uint) (data []model.UserMachineShow, err error) {
	return data, db.DB.Model(model.UserMachineShow{}).
		Preload("MachineInfo").
		Where("user_id =?", userId).
		Where(where).
		Find(&data).Error
}

func EditUserMachine(update model.UserMachine, userId uint) int64 {
	return db.DB.Model(model.UserMachine{}).Where("user_id =?", userId).Updates(update).RowsAffected
}
