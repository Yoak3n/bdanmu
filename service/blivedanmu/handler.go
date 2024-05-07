package blivedanmu

import (
	"bdanmu/api/router/ws"
	"bdanmu/config"
	"bdanmu/consts"
	"bdanmu/package/model"
	"fmt"
	"github.com/Akegarasu/blivedm-go/message"
	"github.com/google/uuid"
	"github.com/tidwall/gjson"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"strings"
)

func messageHandler(msg *message.Danmaku) {
	if msg.Type == message.EmoticonDanmaku {
		msg.Content = fmt.Sprintf("<img src='%s' />", msg.Emoticon.Url)
	} else {
		// 处理emoji表情
		result := gjson.Get(msg.Raw, "info.0.15.extra").String()
		if emots := gjson.Get(result, "emots"); emots.Exists() {
			for k, emot := range emots.Map() {
				width := emot.Get("width").String()
				height := emot.Get("height").String()
				src := emot.Get("url").String()
				msg.Content = strings.ReplaceAll(msg.Content, k, fmt.Sprintf("<img width='%s' src='%s' height='%s' />", width, src, height))
			}
		}
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

	//go func(u *model.User) {
	//	count := 0
	//	for {
	//		reply := <-Queue.Reply
	//		userData, ok := reply[uid]
	//		if !ok && userData == nil {
	//			Queue.Reply <- reply
	//			if count > 10 {
	//				log.Println("用户信息更新失败", u.Name)
	//				break
	//			}
	//			count += 1
	//		} else {
	//			// 成功获得额外用户信息
	//			u.Avatar = userData.Avatar
	//			u.Sex = userData.Sex
	//			u.FollowerCount = userData.FollowerCount
	//			log.Println("用户信息更新成功", u.Name, u.Avatar)
	//			go ws.UpdateUser(u)
	//			break
	//		}
	//		time.Sleep(time.Second)
	//	}
	//}(user)

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

func userEntryHandler(s string) {
	user := NewUserInformation(s)
	m := &model.Message{
		Type: consts.UserEntry,
		Data: user,
	}
	ws.WriteMessage(m)
}
