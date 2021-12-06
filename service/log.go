package service

import (
	"fmt"
	"order/db"
	"order/model"
	"strconv"
)

func CreateLog(orderId int64, text string, num int, status string) {
	log := model.Log{
		OrderID: orderId,
		Text:    text,
		Num:     num,
		Status:  1, // 默认执行中
	}
	if status != "" {
		log.Status, _ = strconv.Atoi(status)
	}
	if err := db.DB.Model(model.Log{}).Create(&log).Error; err != nil {
		fmt.Println(err.Error())
	}
}
