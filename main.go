package main

import (
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
	go initBackend()
	appRun()
}

func appRun() {
	app := NewApp()
	err := wails.Run(&options.App{
		Title:  "",
		Width:  512,
		Height: 1000,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		Frameless:         true,
		HideWindowOnClose: true,
		BackgroundColour:  &options.RGBA{R: 255, G: 255, B: 255, A: 1},
		OnStartup:         app.startup,
		Bind: []interface{}{
			app,
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

func initBackend() {
	util.CreateDirNotExists("data/webview")
	go blivedanmu.Start()
	database.InitDatabase()
	blivedanmu.InitHub()
	c := blivedanmu.GetClient()
	err := c.Start()
	if err != nil {
		logger.Logger.Errorln(err)
	}

}
