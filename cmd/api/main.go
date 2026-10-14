package main

import (
	"heating-prediction/internal/app"
	"heating-prediction/internal/config"
)

func main() {
	cfg := config.LoadFromEnv()

	app := app.NewApp(cfg)

	app.Start()
}
