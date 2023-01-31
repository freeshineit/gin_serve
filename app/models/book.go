package models

type Book struct {
	ID          uint   `json:"id" form:'id' gorm:"primary_key:auto_increment"`
	Title       string `json: "title" form:"title" gorm:"type:varchar(255)"`
	Description string `json:"description" form:"Description" gorm:"type:text"`
	UserID      uint64 `json:"-" form:"-" gorm:"not null"`
	User        User   `json:"user" form:"user" gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE"`
}
