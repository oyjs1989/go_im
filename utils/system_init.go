package utils

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitConfig() {
	// ...
	viper.SetConfigName("app")    // name of config file (without extension)
	viper.AddConfigPath("config") // optionally look for config in the working directory
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error reading config file, %s", err)
	}
	fmt.Println("Config file loaded successfully", viper.Get("app"))
	fmt.Println("Config mysql", viper.Get("db"))
}

func InitMySQL() {
	// ...
	fmt.Println("Init MySQL")
	db, _ = gorm.Open(sqlite.Open(viper.GetString("db")), &gorm.Config{})
}

func GetDB() *gorm.DB {
	return db
}
