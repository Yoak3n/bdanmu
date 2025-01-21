package database

import (
	"bdanmu/internal/model"
	"bdanmu/package/logger"
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
