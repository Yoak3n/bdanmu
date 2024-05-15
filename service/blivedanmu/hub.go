package blivedanmu

import (
	"bdanmu/app"
	"bdanmu/config"
	"bdanmu/package/logger"
	"context"
	"sync"

	"github.com/Akegarasu/blivedm-go/client"
	"github.com/Akegarasu/blivedm-go/message"
)

var (
	cl     *client.Client
	log    = logger.Logger
	mux    = sync.Mutex{}
	appCtx context.Context
)

func InitHub() {
	cl = client.NewClient(config.Conf.RoomId)
	log.Println("connecting room:", cl.RoomID)
	cl.SetCookie(config.Conf.Auth.Cookie)
	var err error
	RoomInfo, err = getRoomInfo(config.Conf.RoomId)
	if err != nil {
		log.Panic(err)
	}
	RegisterHandler()
	appCtx = app.GetApp().Ctx
}

func RegisterHandler() {
	cl.OnDanmaku(messageHandler)
	cl.RegisterCustomEventHandler("INTERACT_WORD", userEntryHandler)
	cl.OnSuperChat(func(s *message.SuperChat) {
		log.Println(s.Message)
	})

}
func GetClient() *client.Client {
	return cl
}
