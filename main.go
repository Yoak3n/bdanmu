package main

import (
	"bdanmu/api/method"
	"bdanmu/app"
	"bdanmu/config"
	"bdanmu/database"
	"bdanmu/package/logger"
	"bdanmu/package/util"
	"bdanmu/service/blivedanmu"
	"embed"
	"fmt"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	go func() {
		util.CreateDirNotExists("data/webview")
		database.InitDatabase()
	}()
	appRun()
}

func appRun() {
	a := app.NewApp()
	blivedanmu.Start()
	go waitForBackendStart(a)
	err := wails.Run(&options.App{
		Title:  "",
		Width:  512,
		Height: 900,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		Frameless:        true,
		DisableResize:    true,
		BackgroundColour: &options.RGBA{R: 28, G: 28, B: 28, A: 1},
		OnStartup:        a.Startup,
		Bind: []interface{}{
			a,
		},
		Windows: &windows.Options{
			DisableFramelessWindowDecorations: false,
			//DisableWindowIcon:   true,
			WebviewUserDataPath: "data/webview",
		},
	})
	if err != nil {
		logger.Logger.Errorln(err)
	}
}

// 由于初期架构问题，只能在主函数中注册事件才能暂时避免循环引用，丑陋的写法
func waitForBackendStart(app *app.App) {
	for {
		if app.Ctx != nil {
			fmt.Println("backend started")
			runtime.EventsOnce(app.Ctx, "start", func(optionalData ...interface{}) {
				method.InitBackend()
				runtime.WindowSetTitle(app.Ctx, blivedanmu.RoomInfo.Title)
				runtime.EventsEmit(app.Ctx, "started", blivedanmu.RoomInfo)
			})
			runtime.EventsOn(app.Ctx, "change", func(id ...interface{}) {
				config.SetRoomId(int(id[0].(float64)))
				method.ChangeBackend()
				runtime.WindowSetTitle(app.Ctx, blivedanmu.RoomInfo.Title)
				runtime.EventsEmit(app.Ctx, "started", blivedanmu.RoomInfo)
			})
			return
		}

	}

}
