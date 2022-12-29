package dao

import (
	"GPT/config"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

// 初始化数据库 Initdatabase
func Initdatabase() {

	cfg := config.Config()

	DSN := fmt.Sprintf("%s:%s@tcp(%s:%d)/db?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.GetString("DatabaseCfg.DriverName.User"),
		cfg.GetString("DatabaseCfg.DriverName.Passwd"),
		cfg.GetString("DatabaseCfg.DriverName.Host"),
		cfg.GetInt("DatabaseCfg.DriverName.Port"),
	)

	// 配置数据库
	DB, err = gorm.Open(mysql.New(
		mysql.Config{
			DriverName:              cfg.GetString("DatabaseCfg.DriverName"),
			DSN:                     DSN,
			DontSupportRenameIndex:  cfg.GetBool("DatabaseCfg.DontSupportRenameIndex"),
			DontSupportRenameColumn: cfg.GetBool("DatabaseCfg.DontSupportRenameColumn"),
		}), &gorm.Config{})

	if err != nil {
		log.Printf("数据库连接失败:%s", err)
	}
}

func GlobalDB() *gorm.DB {
	return DB
}
