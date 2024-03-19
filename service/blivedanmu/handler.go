package blivedanmu

import (
	"bdanmu/api/router/ws"
	"bdanmu/config"
	model2 "bdanmu/package/model"
	"encoding/json"
	"github.com/Akegarasu/blivedm-go/message"
)

const (
	SUPER_CHAT = iota
	UserEntry
	Danmu
)

func messageHandler(msg *message.Danmaku) {
	if msg.Type == message.EmoticonDanmaku {
		log.Println("收到表情包消息:", msg.Emoticon.Url)
	} else {
		danMu := &model2.DanMu{
			Content: msg.Content,
			Sender:  msg.Sender.Uid,
			RoomId:  config.Conf.RoomId,
		}
		m := &model2.Message{
			Type: Danmu,
			Data: danMu,
		}
		data, err := json.Marshal(m)
		if err != nil {
			return
		}
		ws.WriteMessage(data)
		//log.Printf("[弹幕] %s：%s", msg.Sender.Uname, msg.Content)
	}
}

func userEntryHandler(s string) {
	user := NewUserInformation(s)
	m := &model2.Message{
		Type: UserEntry,
		Data: user,
	}
	data, err := json.Marshal(m)
	if err != nil {
		return
	}
	ws.WriteMessage(data)
}
