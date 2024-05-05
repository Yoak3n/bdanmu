package blivedanmu

import (
	"bdanmu/api/router/ws"
	"bdanmu/config"
	"bdanmu/consts"
	"bdanmu/package/model"
	"github.com/tidwall/gjson"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"strings"
	"time"

	"github.com/Akegarasu/blivedm-go/message"
	"github.com/google/uuid"
)

func messageHandler(msg *message.Danmaku) {
	if msg.Type == message.EmoticonDanmaku {
		log.Println("收到表情包消息:", msg.Emoticon.Url)
	} else {
		result := gjson.Get(msg.Raw, "info.0.15.extra").String()

		if emots := gjson.Get(result, "emots"); emots.Exists() {
			for k, emot := range emots.Map() {
				width := emot.Get("width").String()
				height := emot.Get("height").String()
				src := emot.Get("url").String()
				msg.Content = strings.ReplaceAll(msg.Content, k, "<img width='"+width+"' src='"+src+"' height='"+height+"' />")
			}
			log.Println("收到弹幕消息:", msg.Content)
		}

		uid := int64(msg.Sender.Uid)
		user := &model.User{
			UID:   uid,
			Name:  msg.Sender.Uname,
			Guard: msg.Sender.GuardLevel > 0,
		}
		if msg.Sender.Medal != nil {
			user.Medal = &model.Medal{
				Name:     msg.Sender.Medal.Name,
				Level:    msg.Sender.Medal.Level,
				Color:    msg.Sender.Medal.Color,
				TargetID: int64(msg.Sender.Medal.UpUid),
				OwnerID:  uid,
			}
		}

		go SendUserMsg(uid)

		go func(u *model.User) {
			for {
				reply := <-Queue.Reply
				userData, ok := reply[uid]
				if !ok {
					Queue.Reply <- reply
				} else {
					// 成功获得额外用户信息
					if userData == nil {
						break
					}
					u.Avatar = userData.Avatar
					u.Sex = userData.Sex
					u.FollowerCount = userData.FollowerCount
					go ws.UpdateUser(u)
					break
				}
				time.Sleep(time.Second)
			}
		}(user)

		danMu := &model.DanMu{
			Content:   msg.Content,
			Sender:    *user,
			MessageId: uuid.NewString(),
			RoomId:    config.Conf.RoomId,
		}
		go runtime.EventsEmit(appCtx, "danmu", danMu)
		m := &model.Message{
			Type: consts.Danmu,
			Data: danMu,
		}

		go ws.WriteMessage(m)
	}
}

func userEntryHandler(s string) {
	user := NewUserInformation(s)
	m := &model.Message{
		Type: consts.UserEntry,
		Data: user,
	}
	ws.WriteMessage(m)
}
