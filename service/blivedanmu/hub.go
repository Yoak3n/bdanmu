package blivedanmu

import (
	"bdanmu/config"
	"bdanmu/package/logger"
	"github.com/Akegarasu/blivedm-go/client"
	"github.com/tidwall/gjson"
	"sync"
)

var (
	cl  *client.Client
	log = logger.Logger
	mux = sync.Mutex{}
)

func InitHub() {
	cl = client.NewClient(config.Conf.RoomId)
	cl.SetCookie(config.Conf.Auth.Cookie)
	var err error
	roomInfo, err = getRoomInfo(config.Conf.RoomId)
	if err != nil {
		log.Panic(err)
	}
	RegisterHandler()
}

func RegisterHandler() {
	cl.OnDanmaku(messageHandler)
	cl.RegisterCustomEventHandler("WELCOME", func(s string) {
		data := gjson.Get(s, "data").String()
		log.Println("欢迎", data)
	})
	cl.RegisterCustomEventHandler("INTERACT_WORD", userEntryHandler)
	cl.RegisterCustomEventHandler("SUPER_CHAT_MESSAGE", func(s string) {
		result := gjson.Parse(s)
		log.Println(result.Get("data").String())
	})
}
func GetClient() *client.Client {
	return cl
}
