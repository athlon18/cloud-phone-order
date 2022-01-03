package model

type Machine struct {
	GlobalModel
	MachineCode string `json:"machine_code" gorm:"index"` //机器码
	//Status      bool   `json:"status"`                    // 状态
}
