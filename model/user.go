package model

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type User struct {
	gorm.Model
	Username string  `json:"username"`
	Email    string  `json:"email" valid:"email" gorm:"unique"`
	Password string  `json:"password" valid:"length(6|16)"`
	Photos   []Photo `json:"photos" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
