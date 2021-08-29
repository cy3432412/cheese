package mysql

import (
	"cheese/setting"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Init(cfg *setting.MySQLConfig) (err error) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)
	db, err = gorm.Open("mysql", dsn)
	db.SingularTable(true)
	if err != nil {
		return
	}
	sqlDB := db.DB()
	// SetMaxIdleCons 设置连接池中的最大闲置连接数。
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)

	// SetMaxOpenCons 设置数据库的最大连接数量。
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)

	return db.DB().Ping()
}

func Close() {
	_ = db.Close()
}
