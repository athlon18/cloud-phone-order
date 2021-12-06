package model

type File struct {
	GlobalModel
	Name   string `gorm:"name" json:"name"` // 文件名称
	File   string `gorm:"file" json:"file"` // 文件编码
	Status bool   `gorm:"status" json:"-"`  // 状态
}
