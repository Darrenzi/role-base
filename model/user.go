package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `form:"username" gorm:"type:varchar(32)" validate:"required"`
	Password string `form:"password" gorm:"type:varchar(32)" validate:"required"`
	Salt     string `gorm:"type:char(32)"`
	Roles    []Role `gorm:"many2many:user_role;"`
}

type LoginForm struct {
	Username string `form:"username" validate:"required"`
	Password string `form:"password" validate:"required"`
}
