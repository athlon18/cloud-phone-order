package model

type Log struct {
	CreatedModel
	OrderID int64  `gorm:"order_id" json:"order_id"` // 订单ID
	Text    string `gorm:"text" json:"text"`         // 内容
	Num     int    `gorm:"num" json:"num"`           // 数量
	Status  int    `gorm:"status" json:"status"`     // 状态
}
