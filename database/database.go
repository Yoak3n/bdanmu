package database

import (
	"bdanmu/config"
	"gorm.io/gorm"
	"time"
)

var (
	db *gorm.DB
)

func InitDatabase() {
	switch config.Conf.Database.Type {
	case "postgres":
		db = initPostgres(config.Conf.Database.Host, config.Conf.Database.User, config.Conf.Database.Password, config.Conf.Database.Name, config.Conf.Database.Port)
	case "sqlite", "sqlite3":
		db = initSqlite(config.Conf.Database.Name)
	default:
		panic("Unsupported database type,please check the configuration")
	}
	if db == nil {
		return
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&UserTable{})
	if err != nil {
		panic(err)
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)
	log.Println("database connected")
}

func GetDB() *gorm.DB {
	return db
}
