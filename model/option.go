package model

type Option struct {
	GlobalModel
	Name       string       `gorm:"name" json:"name"`                                   // 选项名称
	Code       string       `gorm:"code" json:"code"`                                   // 选项编码
	GameId     string       `gorm:"game_id" json:"game_id"`                             // 游戏ID
	OptionItem []OptionItem `gorm:"foreignkey:option_id" json:"option_items,omitempty"` // 选项内容
}

type OptionItem struct {
	GlobalModel
	Name     string `gorm:"name" json:"name"`           // 选项内容名称
	Code     string `gorm:"code" json:"code"`           // 选项内容编码
	OptionId uint   `gorm:"option_id" json:"option_id"` // 选项ID
}
