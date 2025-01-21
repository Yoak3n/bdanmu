package service

import (
	"bdanmu/database"
	"bdanmu/internal/model"
	"gorm.io/gorm/clause"
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
		database.GetDB().Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "uid"}},
			DoUpdates: clause.AssignmentColumns([]string{"updated_at", "name", "sex", "avatar", "follower_count", "medal_name", "medal_owner_id", "medal_level", "medal_target_id", "guard", "medal_color"}),
		}).Create(record)
	}
	mux.Unlock()
}

func CreateUserAndUpdateStack(users []*model.User) {
	records := make([]database.UserTable, 0)
	for _, user := range users {
		record := &database.UserTable{}
		record.User = *user
		records = append(records, *record)
	}
	database.GetDB().Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "uid"}},
		DoUpdates: clause.AssignmentColumns([]string{"updated_at", "name", "sex", "avatar", "follower_count", "medal_name", "medal_owner_id", "medal_level", "medal_target_id", "guard", "medal_color"}),
	}).Create(records)
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
