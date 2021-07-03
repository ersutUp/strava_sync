package db

import (
	"fit_sync_server/conf/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)


var Mydb *ExtendDB

//连接sqlite数据库
func Connect() {
	db, err := gorm.Open(sqlite.Open("./fit_sync.db"), &gorm.Config{
		Logger: &log.Gorm2logrus{
			SlowThreshold: 5*time.Second,
		},
	})
	if err != nil {
		panic("failed to connect database")
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic("failed sqlDB")
	}

	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(3)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(30)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(10*time.Second)

	Mydb = &ExtendDB{DB: db}
}