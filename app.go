package main

import (
	"bdanmu/config"
	"context"
)

// App struct
type App struct {
	ctx context.Context
}
type Result struct {
	Name   string `json:"name"`
	Status bool   `json:"status"`
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) LoginSubmit() {
	config.AuthBilibili()
}
