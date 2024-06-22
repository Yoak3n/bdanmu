package blivedanmu

import (
	"bdanmu/app"
	"bdanmu/config"
	"bdanmu/package/logger"
	"context"
	"sync"

	"github.com/Akegarasu/blivedm-go/client"
)

var (
	cl     *client.Client
	log    = logger.Logger
	mux    = sync.Mutex{}
	appCtx context.Context
)

func InitHub() (err error) {
	RoomInfo, err = getRoomInfo(config.Conf.RoomId)
	if err != nil {
		return err
	}
	cl = client.NewClient(config.Conf.RoomId)
	log.Println("connecting room:", cl.RoomID)
	cl.SetCookie(config.Conf.Auth.Cookie)

	RegisterHandler()
	appCtx = app.GetApp().Ctx
	return nil
}

func RegisterHandler() {
	cl.OnDanmaku(messageHandler)
	cl.RegisterCustomEventHandler("INTERACT_WORD", userEntryHandler)
	cl.OnSuperChat(superChatHandler)

}
func GetClient() *client.Client {
	return cl
}
