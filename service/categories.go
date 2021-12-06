package service

import (
	"gorm.io/gorm"
	"order/common/util"
	"order/db"
	"order/model"
	"strconv"
	"strings"
	"time"
)

type Categories struct {
	Value    string        `json:"value"`
	Label    string        `json:"label"`
	Num      int           `json:"num"`
	Rules    string        `json:"rules"`
	Hours    string        `json:"hours"`
	Children []*Categories `json:"children,omitempty"`
}

func GetCategoriesList(gameId int) []*Categories {
	return CategoriesList(uint(gameId), 0)
}

func CategoriesList(gameId uint, id uint) []*Categories {

	var categories []*Categories
	for _, item := range getCategoriesList(db.DB, gameId, id) {
		categories = append(categories, &Categories{
			Label:    item.Name,
			Value:    strconv.Itoa(int(item.ID)),
			Num:      item.Num,
			Rules:    item.Rules,
			Hours:    item.Hours,
			Children: CategoriesList(gameId, item.ID),
		})
	}
	return categories
}

func getCategoriesList(db *gorm.DB, gameId uint, id uint) []model.Categories {
	var data []model.Categories
	if err := db.Model(model.Categories{}).Where("game_id = ? and parent = ? and status = 1", gameId, id).Find(&data).Error; err != nil {
		return nil
	}
	var list []model.Categories
	for _, item := range data {
		if item.Rules != "" {
			arr := strings.Split(item.Rules, ",")
			sd, _ := time.ParseDuration("-" + item.Hours + "h")
			week := int(time.Now().Add(sd).Weekday())
			if week == 0 { // 处理周日为0的问题
				week = 7
			}
			if util.Contains(arr, strconv.Itoa(week)) == true {
				list = append(list, item)
			}
		} else {
			list = append(list, item) // 默认为空 就全开启
		}
	}
	return list
}

func GetCategoryName(types string) string {
	arr := strings.Split(types, ",")
	var data []model.Categories
	if err := db.DB.Model(model.Categories{}).Where("id in ? ", arr).Find(&data).Error; err != nil {

	}
	var str []string
	for _, item := range data {
		str = append(str, item.Name)
	}
	return strings.Join(str, "/")
}

func GetCategoryCode(types string) string {
	arr := strings.Split(types, ",")
	var data []model.Categories
	if err := db.DB.Model(model.Categories{}).Where("id in ? ", arr).Find(&data).Error; err != nil {

	}
	var str []string
	for _, item := range data {
		str = append(str, item.Code)
	}
	return strings.Join(str, "_")
}
