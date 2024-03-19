package database

import (
	"bdanmu/package/logger"
	model2 "bdanmu/package/model"
	"gorm.io/gorm"
)

var log = logger.Logger

type UserTable struct {
	gorm.Model
	model2.User
}

type DanMuTable struct {
	gorm.Model
	model2.DanMu
}
