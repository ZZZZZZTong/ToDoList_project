package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	UserName string `gorm:"unique"`
	PasswordDigest string //密码存储的是密文
}
