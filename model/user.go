package model

type User struct {
	GlobalModel
	Name     string `gorm:"name" json:"name"`                             // 用户账号
	Password string `gorm:"->:false;<-:create" json:"password,omitempty"` // 用户密码
	Code     string `gorm:"code" json:"code"`                             // 特征码
	Token    string `gorm:"token" json:"token"`                           // 登录密钥
	Machine  string `gorm:"machine" json:"machine"`                       // 机器码
}

type UserMachine struct {
	GlobalModel
	UserId    uint   `json:"user_id" gorm:"index"`     // 用户id
	Machine   string `json:"machine" gorm:"unique"`    // 用户机器码
	MachineId uint   `json:"machine_id" gorm:"unique"` // 机器码ID
	Tag       string `json:"tag"`                      // 标签
}
