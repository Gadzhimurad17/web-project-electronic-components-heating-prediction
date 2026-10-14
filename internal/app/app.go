package app

import (
	"context"
	f "heating-prediction/internal/app/factory"
	"heating-prediction/internal/config"
	"heating-prediction/internal/interfaces/http"
)

type App struct {
	server *http.Server
}

func NewApp(cfg *config.Config) *App {
	factory := f.NewFactory(cfg)
	router := factory.NewRouter()

	handlers := factory.NewHandlers()

	server := factory.NewHttpServer(
		cfg.DbHost,
		cfg.AppPort,
		handlers,
		router,
	)

	return &App{
		server: server,
	}
}

func (app *App) Start() {
	app.server.Start()
}

func (app *App) Stop(ctx context.Context) {
	app.server.Close(ctx)
}
