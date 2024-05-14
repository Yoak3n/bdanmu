package main

import (
	"bdanmu/api/method"
	"bdanmu/app"
	"bdanmu/database"
	"bdanmu/package/logger"
	"bdanmu/package/util"
	"bdanmu/service/blivedanmu"
	"embed"
	"time"

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
	go backendStart(a)
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

func backendStart(app *app.App) {

	for {
		time.Sleep(1 * time.Second)
		if app.Ctx != nil {
			runtime.EventsOnce(app.Ctx, "start", func(optionalData ...interface{}) {
				method.InitBackend()
				runtime.WindowSetTitle(app.Ctx, blivedanmu.RoomInfo.Title)
				runtime.EventsEmit(app.Ctx, "started", blivedanmu.RoomInfo)
			})
			return
		}

	}

}
