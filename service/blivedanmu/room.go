package blivedanmu

import (
	"bdanmu/config"
	"bdanmu/internal/model"
	"bdanmu/package/request"
	"errors"
	"fmt"
	"github.com/tidwall/gjson"
)

var RoomInfo *model.Room

func getRoomInfo(id int) (*model.Room, error) {
	res, err := request.Get("https://api.live.bilibili.com/room/v1/Room/get_info", fmt.Sprintf("room_id=%d", config.Conf.RoomId))
	if err != nil {
		return nil, err
	}
	room := &model.Room{
		ShortId: id,
		User:    model.User{},
	}
	result := gjson.ParseBytes(res)
	if result.Get("code").Int() == 0 {
		room.User.UID = result.Get("data.uid").Int()
		room.LongId = result.Get("data.room_id").Int()
		room.FollowerCount = result.Get("data.attention").Int()
		room.Title = result.Get("data.title").String()
		room.Cover = result.Get("data.user_cover").String()
		user := getUserInfo(room.User.UID)
		if user != nil {
			room.User = *user
		}
		return room, nil
	}
	return nil, errors.New("get room information failed")

}
