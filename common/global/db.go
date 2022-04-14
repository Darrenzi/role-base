package global

import (
	"blog/model"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var _db *gorm.DB

func InitDb() {
	conf := Config.Database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", conf.Username, conf.Password, conf.Host, conf.Port, conf.Name)
	var err error
	_db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic(fmt.Sprintf("连接数据库失败: %v\n", err.Error()))
	}

	sqlDB, _ := _db.DB()

	//设置数据库连接池最大连接数
	sqlDB.SetMaxOpenConns(conf.MaxOpenConns)
	//连接池最大允许的空闲连接数
	sqlDB.SetMaxIdleConns(conf.MaxIdleConns)

	_db.AutoMigrate(&model.Role{}, &model.User{})
}

func GetDB() *gorm.DB {
	return _db
}
