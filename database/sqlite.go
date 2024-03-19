package database

import (
	"bdanmu/package/logger"
	"bdanmu/package/util"
	"fmt"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func initSqlite(name string) *gorm.DB {
	util.CreateDirNotExists("data/db")
	dsn := fmt.Sprintf(fmt.Sprintf("data/db/%s.db", name))
	sdb, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Logger.Println(fmt.Sprintf("database connected err:%v", err))
	}
	if err != nil {
		logger.Logger.Println(fmt.Sprintf("database connected err:%v", err))
	}
	return sdb
}
