package models

import (
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func TestMain(m *testing.M) {
	var err error
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	db.AutoMigrate(&UserBasic{})

	m.Run()
}

func TestDBOperations(t *testing.T) {
	// Create
	db.Create(&UserBasic{Name: "D42", PassWord: "100"})

	// Read
	var user UserBasic
	db.First(&user, 1)                 // 根据整型主键查找
	db.First(&user, "Name = ?", "D42") // 查找 code 字段值为 D42 的记录

	if user.Name != "D42" {
		t.Errorf("Expected Name to be D42, got %s", user.Name)
	}

	// Update - 将 product 的 price 更新为 200
	db.Model(&user).Update("Price", 200)
	// Update - 更新多个字段
	db.Model(&user).Updates(UserBasic{Name: "200", Phone: "F42"}) // 仅更新非零值字段
	db.Model(&user).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - 删除 product
	db.Delete(&user, 1)
}
