package service

import (
	"order/db"
	"order/model"
	"time"
)

func ExternalRegister(machineCode string) error {
	return db.DB.Model(model.Machine{}).Assign(model.Machine{MachineCode: machineCode}).FirstOrCreate(&model.Machine{MachineCode: machineCode}).Error
}

func ExternalDeregister(machineCode string) (int64, error) {
	data := db.DB.Model(model.Machine{}).Where("machine_code =?", machineCode).Delete(new(model.Machine))
	return data.RowsAffected, data.Error
}

func ExternalHealth(machineCode string) (int64, error) {
	data := db.DB.Model(model.Machine{}).Where("machine_code =?", machineCode).Update("updated_at", time.Now())
	return data.RowsAffected, data.Error
}
