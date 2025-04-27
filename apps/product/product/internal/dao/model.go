package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"tiny_service/apps/product/product/internal/config"
)

var gdb *gorm.DB

func InitGDB(conf *config.MysqlConf) {
	var err error
	gdb, err = gorm.Open(mysql.Open(conf.DataSource), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}
