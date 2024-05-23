package app

import (
	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func registerTray() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetTemplateIcon(icon.Data, icon.Data)
	systray.SetTitle("Bdanmu")
	systray.SetTooltip("Bdanmu")
	mShow := systray.AddMenuItem("显示", "")
	mQuit := systray.AddMenuItem("退出", "Quit the whole app")
	go func() {
		for {
			select {
			case <-mShow.ClickedCh:
				runtime.WindowShow(app.Ctx)
				systray.Quit()
			case <-mQuit.ClickedCh:
				runtime.Quit(app.Ctx)
			}
		}
	}()
}

func onExit() {
}
