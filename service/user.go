package service

import (
	"bdanmu/database"
	"bdanmu/package/model"
	"sync"
	"time"
)

var mux sync.Mutex

func CreateUserRecord(user *model.User) {

	record := &database.UserTable{}
	record.User = *user
	mux.Lock()
	if result := database.GetDB().Where("uid = ?", user.UID).Find(record); result.RowsAffected > 0 && time.Since(record.UpdatedAt).Hours() > 24*5 {
		UpdateUserRecord(user)
	} else if result.RowsAffected <= 0 {
		database.GetDB().Create(record)
	}
	mux.Unlock()
}

func ReadUserRecord(uid int64) *model.User {
	record := &database.UserTable{}
	if result := database.GetDB().Where("uid = ?", uid).Find(record); result.RowsAffected > 0 && time.Since(record.UpdatedAt).Hours() < 24*30 {
		return &record.User
	}
	return nil
}

func UpdateUserRecord(user *model.User) {
	record := &database.UserTable{}
	record.User = *user
	database.GetDB().Model(&database.UserTable{}).Where("uid = ? ", user.UID).Omit("uid", "id").Updates(record)
}

//func DeleteUserRecord(uid int64) {
//}
