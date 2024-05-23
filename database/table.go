package database

import (
	"bdanmu/package/logger"
	"bdanmu/package/model"
	"gorm.io/gorm"
)

var log = logger.Logger

type UserTable struct {
	gorm.Model
	model.User
}

type DanMuTable struct {
	gorm.Model
	model.DanMu
}
