package models

import (
	"encoding/json"
	"fmt"
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
	jsonStr := `{"id":1,"name":"D42","passWord":"100"}`
	dto := UserDTO{}
	json.Unmarshal([]byte(jsonStr), &dto)
	mStr, _ := json.Marshal(dto)
	fmt.Printf("mStr: %v\n", string(mStr))
	m := UserBasic{}
	json.Unmarshal(mStr, &m)
	fmt.Printf("m: %+v\n", m)
	// 清除表数据
	db.Exec("DELETE FROM user_basics")
	// Save
	db.Save(&m)

	// Read
	var user UserBasic
	db.First(&user, "Name = ?", "D42") // 查找 code 字段值为 D42 的记录
	fmt.Printf("user: %+v\n", user)
}
