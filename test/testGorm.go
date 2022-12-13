package main

import (
	"im_go/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var mysqlconfig = "root:123456"

func main() {
	db, err := gorm.Open(mysql.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	db.AutoMigrate(&models.UserBasic{})

	// Create
	db.Create(&models.UserBasic{Name: "D42", PassWord: "100"})

	// Read
	var user models.UserBasic
	db.First(&user, 1)                 // 根据整型主键查找
	db.First(&user, "Name = ?", "D42") // 查找 code 字段值为 D42 的记录

	// Update - 将 product 的 price 更新为 200
	db.Model(&user).Update("Price", 200)
	// Update - 更新多个字段
	db.Model(&user).Updates(models.UserBasic{Name: "200", Phone: "F42"}) // 仅更新非零值字段
	db.Model(&user).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - 删除 product
	db.Delete(&user, 1)
}
