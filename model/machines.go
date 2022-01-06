package model

type Machine struct {
	GlobalModel
	MachineCode string `json:"machine_code" gorm:"index"` //机器码
	//Status      bool   `json:"status"`                    // 状态
}

type MachineShow struct {
	ShowModel
	MachineCode string `json:"-" gorm:"index"` //机器码
}

func (MachineShow) TableName() string {
	return "machines"
}
