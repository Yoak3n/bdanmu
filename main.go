package main

import (
	"bdanmu/api/router"
	"bdanmu/app"
	"bdanmu/app/runtime"

	"bdanmu/database"
	"bdanmu/package/logger"
	"bdanmu/package/util"
	"bdanmu/service/blivedanmu"
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	go func() {
		util.CreateDirNotExists("data/webview")
		database.InitDatabase()
		router.InitRouter()
	}()
	appRun()
}

func appRun() {
	a := app.NewApp()
	blivedanmu.Start()
	// register global event(把丑陋的东西写在看起来正常的地方那它就看起来正常了)
	go func(app *app.App) {
		for {
			if app.Ctx != nil {
				runtime.RegisterSetRoomId(&app.Ctx)
				return
			}
		}
	}(a)
	err := wails.Run(&options.App{
		Title:  "",
		Width:  512,
		Height: 900,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		Frameless:        true,
		MaxWidth:         512,
		MinWidth:         512,
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
