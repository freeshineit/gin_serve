package model

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	ID          uint   `json:"id" form:"id" gorm:"primary_key:auto_increment"`
	Title       string `json:"title" form:"title" gorm:"type:varchar(255)"`
	Description string `json:"description" form:"description" gorm:"type:text"`
	UserID      uint64 `json:"-" form:"-" gorm:"not null"`
	User        User   `json:"user" form:"user" gorm:"foreign_key:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE"`
}
