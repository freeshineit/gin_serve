package model

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	ID        uint           `json:"id" gorm:"primary_key"`
	CreatedAt time.Time      `json:"create_at" gorm:"column:create_at"`
	UpdatedAt time.Time      `json:"update_at" gorm:"column:update_at"`
	DeletedAt gorm.DeletedAt `json:"delete_at" gorm:"index;column:update_at"`
}
