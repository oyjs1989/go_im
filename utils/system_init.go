package utils

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
	// 自定义打印sql
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, // 满sql
			LogLevel:      logger.Info,
			Colorful:      true, //彩色
		},
	)

	// ...
	fmt.Println("Init MySQL")
	db, _ = gorm.Open(sqlite.Open(viper.GetString("db")), &gorm.Config{Logger: newLogger})
}

func GetDB() *gorm.DB {
	return db
}
