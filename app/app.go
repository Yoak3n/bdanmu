package app

import (
	"bdanmu/config"
	"context"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	Ctx context.Context
}
type Result struct {
	Name   string `json:"name"`
	Status bool   `json:"status"`
}

var app *App

// NewApp creates a new App application struct
func NewApp() *App {
	app = &App{}
	return app
}

func GetApp() *App {
	return app
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.Ctx = ctx
}

func (a *App) NeedLogin(cookie string) bool {
	logged, _, cookieStr, csrf := config.IsLogin(cookie)
	if logged {
		config.Conf.Auth.Cookie = cookieStr
		config.Conf.Auth.RefreshToken = csrf
		config.SetCookieRefresh()
		runtime.EventsEmit(app.Ctx, "auth", cookieStr, csrf)
		return false
	} else {
		return true
	}
}

func (a *App) LoginBilibili() string {
	loginUrl, loginKey := config.LoginFromFrontend()
	go waitScanToLogin(loginKey)
	return loginUrl
}

func waitScanToLogin(key string) {
	config.VerifyLogin(key)
	logged, _, cookieStr, csrf := config.IsLogin()
	if logged {
		runtime.EventsEmit(app.Ctx, "auth", cookieStr, csrf)
		config.Conf.Auth.Cookie = cookieStr
		config.Conf.Auth.RefreshToken = csrf
		config.SetCookieRefresh()
	}
}

func (a *App) HideToTray() {
	runtime.Hide(a.Ctx)
	go registerTray()
}

func (a *App) SyncAuth() []string {
	return []string{config.Conf.Auth.Cookie, config.Conf.Auth.RefreshToken} //config.Conf.Auth.Cookie, config.Conf.Auth.RefreshToken
}

func (a *App) SyncRoom() int {
	return config.Conf.RoomId
}
