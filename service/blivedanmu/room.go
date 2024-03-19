package blivedanmu

import (
	"bdanmu/config"
	model2 "bdanmu/package/model"
	"bdanmu/package/request"
	"errors"
	"fmt"
	"github.com/tidwall/gjson"
)

var roomInfo *model2.Room

func getRoomInfo(id int) (*model2.Room, error) {
	res, err := request.Get("https://api.live.bilibili.com/room/v1/Room/get_info", fmt.Sprintf("room_id=%d", config.Conf.RoomId))
	if err != nil {
		return nil, err
	}
	room := &model2.Room{
		ShortId: id,
		User:    model2.User{},
	}
	result := gjson.ParseBytes(res)
	if result.Get("code").Int() == 0 {
		room.User.UID = result.Get("data.uid").Int()
		room.LongId = result.Get("data.room_id").Int()
		room.FollowerCount = result.Get("data.attention").Int()
		user := getUserInfo(room.User.UID)
		if user != nil {
			room.User = *user
		}
		return room, nil
	}
	return nil, errors.New("get room information failed")

}
