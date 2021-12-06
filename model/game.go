package model

type Game struct {
	GlobalModel
	Name   string   `gorm:"name" json:"name"`                              // 游戏名称
	Code   string   `gorm:"code" json:"code"`                              // 游戏编码
	Ver    string   `gorm:"ver" json:"ver"`                                // 版本号
	Status bool     `gorm:"status" json:"-"`                               // 游戏启用状态
	Mode   []Mode   `gorm:"foreignkey:game_id" json:"mode_list,omitempty"` // 游戏模式
	Option []Option `gorm:"foreignkey:game_id" json:"option,omitempty"`    // 游戏选项
}
