package models

import (
	"im_go/utils"
	"time"

	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Name          string
	PassWord      string
	Phone         string
	Email         string
	Identity      string
	ClientIp      string
	ClientPort    string
	LoginTime     time.Time
	HeartBeatTime time.Time
	LoginOutTime  time.Time
	IsLogout      bool
	DeviceInfo    string
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}

func GetOne() *UserBasic {
	user := UserBasic{}
	db := utils.GetDB()
	db.Find(&user)
	return &user
}

func CreateOne(name string, pw string) *UserBasic {
	user := UserBasic{Name: name, PassWord: pw}
	db := utils.GetDB()
	db.Create(&user)
	return &user
}

// 查找用户
func FindUser(name string) *UserBasic {
	user := UserBasic{}
	db := utils.GetDB()
	db.Where("name = ?", name).Find(&user)
	return &user
}
