package runtime

import (
	"bdanmu/api/method"
	"bdanmu/config"
	"bdanmu/package/logger"
	"bdanmu/service/blivedanmu"
	"context"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func RegisterSetRoomId(ctx *context.Context) {
	runtime.EventsOn(*ctx, "change", func(id ...interface{}) {
		config.SetRoomId(int(id[0].(float64)))
		err := method.ChangeBackend()
		if err != nil {
			logger.Logger.Errorln(err)
			runtime.EventsEmit(*ctx, "error", err.Error())
		} else {
			runtime.WindowSetTitle(*ctx, blivedanmu.RoomInfo.Title)
			runtime.EventsEmit(*ctx, "started", blivedanmu.RoomInfo)
		}
	})

}
