package service

import (
	"order/db"
	"order/model"
	"testing"
)

func TestMachines(t *testing.T) {
	if err := db.Conn().AutoMigrate(model.Machine{}, model.UserMachine{}); err != nil {
		return
	}
}
