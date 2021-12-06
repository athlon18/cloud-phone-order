package model

type Categories struct {
	GlobalModel
	GameId uint   `gorm:"game_id" json:"game_id"` // 游戏id
	Name   string `gorm:"name" json:"name"`       // 类目名称
	Code   string `gorm:"code" json:"code"`       // 类目编码
	Parent uint   `gorm:"parent" json:"parent"`   // 上级
	Num    int    `gorm:"num" json:"num"`         // 数量上限
	Rules  string `gorm:"rules" json:"rules"`     // 日期规则
	Hours  string `gorm:"hours" json:"hours"`     // 小时
	Status bool   `gorm:"status" json:"-"`        // 状态
}
