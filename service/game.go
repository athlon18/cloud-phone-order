package service

import (
	"order/db"
	"order/model"
)

func GameList() ([]model.Game, error) {
	var data []model.Game
	if err := db.DB.Model(model.Game{}).
		Preload("Mode").
		Preload("Option").
		Preload("Option.OptionItem").
		Where("status = 1").Find(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}
