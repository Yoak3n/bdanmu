package app

import (
	_ "embed"
	"github.com/getlantern/systray"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed icon.ico
var icon []byte

func registerTray() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetTemplateIcon(icon, icon)
	systray.SetTitle("Bdanmu")
	systray.SetTooltip("Bdanmu")
	mShow := systray.AddMenuItem("显示", "")
	mMinimize := systray.AddMenuItem("最小化", "Minimize the app")
	mQuit := systray.AddMenuItem("退出", "Quit the whole app")
	go func() {
		for {
			select {
			case <-mShow.ClickedCh:
				runtime.WindowShow(app.Ctx)
				systray.Quit()
			case <-mMinimize.ClickedCh:
				runtime.WindowMinimise(app.Ctx)
			case <-mQuit.ClickedCh:
				runtime.Quit(app.Ctx)
			}
		}
	}()
}

func onExit() {
}
