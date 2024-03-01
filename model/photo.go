package model

import (
	"gorm.io/gorm"
)

type Photo struct {
	gorm.Model
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	photoUrl string `json:"photoUrl"`
	UserId   uint   `gorm:"foreignKey:UserID;references:ID"`
}
