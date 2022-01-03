package model

import (
	"gorm.io/gorm"
	"time"
)

type GlobalModel struct {
	ID        uint           `gorm:"AUTO_INCREMENT" json:"id"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type CreatedModel struct {
	ID        uint           `gorm:"AUTO_INCREMENT" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
