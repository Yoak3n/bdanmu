package blivedanmu

import (
	"bdanmu/api/router/ws"
	"bdanmu/config"
	"bdanmu/consts"
	"bdanmu/package/logger"
	"bdanmu/package/model"
	"fmt"
	"strings"

	"github.com/Akegarasu/blivedm-go/message"
	"github.com/google/uuid"
	"github.com/tidwall/gjson"
	"github.com/wailsapp/wails/v2/pkg/runtime"
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
	logger.Logger.Traceln(msg.Content)
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

	danMu := &model.DanMu{
		Content:   msg.Content,
		Sender:    *user,
		MessageId: uuid.NewString(),
		RoomId:    config.Conf.RoomId,
	}
	go runtime.EventsEmit(appCtx, "danmu", danMu)
	m := &model.Message{
		Type: consts.DAN_MU,
		Data: danMu,
	}
	ws.WriteMessage(m)

}

func userEntryHandler(s string) {
	user := NewUserInformation(s)
	m := &model.Message{
		Type: consts.USER_ENTRY,
		Data: user,
	}
	ws.WriteMessage(m)
}

func superChatHandler(s *message.SuperChat) {
	superChat := &model.SuperChat{
		User: model.User{
			UID:    int64(s.Uid),
			Name:   s.UserInfo.Uname,
			Avatar: s.UserInfo.Face,
			Guard:  s.UserInfo.GuardLevel > 0,
			Medal: &model.Medal{
				Name:     s.MedalInfo.MedalName,
				OwnerID:  int64(s.Uid),
				Level:    s.MedalInfo.MedalLevel,
				TargetID: int64(s.MedalInfo.TargetId),
			},
		},
		RoomID:    config.Conf.RoomId,
		MessageID: uuid.NewString(),
		Content:   s.Message,
		Timestamp: s.Ts,
		EndTime:   s.EndTime,
		Price:     s.Price,
	}
	m := &model.Message{
		Type: consts.SUPER_CHAT,
		Data: superChat,
	}
	go runtime.EventsEmit(appCtx, "superChat", superChat)
	ws.WriteMessage(m)
}
