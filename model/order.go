package model

type Order struct {
	GlobalModel
	OrderId  int64  `gorm:"order_id" json:"order_id"`                                     // 订单id
	GameId   uint   `gorm:"game_id" json:"-"`                                             // 游戏id
	ModeId   uint   `gorm:"mode_id" json:"-"`                                             // 模式id
	UserId   uint   `gorm:"user_id" json:"-"`                                             // 用户id
	Name     string `gorm:"name" json:"name"`                                             // 账号用户名
	Password string `gorm:"password" json:"password"`                                     // 账号密码
	Type     string `gorm:"type" json:"type"`                                             // 类型
	Option   string `gorm:"option" json:"option"`                                         // 选项
	Num      int    `gorm:"num" json:"num"`                                               // 数量
	CNum     int    `gorm:"column:cnum" json:"cnum"`                                      // 完成数量
	Status   int    `gorm:"status" json:"status"`                                         // 订单状态 0 初始化 1 执行中，2 执行完毕， -1 执行失败，-2 暂停订单
	Mode     Mode   `gorm:"foreignKey:mode_id;references:id" json:"mode,omitempty"`       // 模式详情
	Game     Game   `gorm:"foreignKey:game_id;references:id" json:"game,omitempty"`       // 游戏详情
	Log      []Log  `gorm:"foreignKey:order_id;references:order_id" json:"log,omitempty"` // 日志详情
	Machine  string `json:"machine" gorm:"index"`                                         // 绑定机器码
}
