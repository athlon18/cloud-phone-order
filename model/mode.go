package model

type Mode struct {
	GlobalModel
	GameId uint   `gorm:"game_id" json:"-"` // 游戏id
	Name   string `gorm:"name" json:"name"` // 模式名称 （例：安卓，ios）
	Code   string `gorm:"code" json:"code"` // 模式编码
}
