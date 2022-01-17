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

func GetUserMachine(userId uint) (data model.UserMachineShow, err error) {
	return data, db.DB.Debug().Raw(
		`select * FROM user_machines where machine = (SELECT u.machine FROM user_machines u 
			LEFT JOIN orders o ON u.machine = o.machine and o.status != 2
			WHERE u.user_id = ?  
			GROUP BY u.machine ORDER BY count(u.machine)  asc LIMIT 1)`, userId).First(&data).Error
}

func EditUserMachine(update model.UserMachine, id string, userId uint) int64 {
	return db.DB.Model(model.UserMachine{}).
		Where("id =?", id).
		Where("user_id =?", userId).
		Updates(update).RowsAffected
}
