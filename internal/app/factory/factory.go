package factory

import (
	"heating-prediction/internal/config"
)

type Factory struct {
	cfg *config.Config
}

func NewFactory(cfg *config.Config) *Factory {
	return &Factory{
		cfg: cfg,
	}
}
