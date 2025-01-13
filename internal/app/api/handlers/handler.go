package handlers

import (
	"github.com/arafetki/go-echo-boilerplate/internal/config"
	"github.com/arafetki/go-echo-boilerplate/internal/logging"
	"github.com/arafetki/go-echo-boilerplate/internal/services"
)

type Handler struct {
	services *services.Service
	config   config.Config
	logger   *logging.Wrapper
}

func New(svc *services.Service, cfg config.Config, logger *logging.Wrapper) *Handler {
	return &Handler{
		services: svc,
		config:   cfg,
		logger:   logger,
	}
}
