package service

import (
	"encoding/json"
	"gorm.io/gorm"
	"order/db"
	"order/model"
)

func GetOptionSelectCodeByIds(str string) []map[string]string {
	var data map[string]int
	_ = json.Unmarshal([]byte(str), &data)
	var result []map[string]string
	for index, item := range data {
		option := model.Option{}
		if err := db.DB.Model(model.Option{}).Preload("OptionItem", func(db *gorm.DB) *gorm.DB {
			return db.Where("id = ?", item)
		}).Where("id = ?", index).First(&option).Error; err != nil {
			continue
		}
		result = append(result, map[string]string{option.Code: option.OptionItem[0].Code})
	}
	return result
}
