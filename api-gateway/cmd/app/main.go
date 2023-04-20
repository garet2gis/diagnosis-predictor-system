package main

import (
	"api-gateway/internal/config"
	"api-gateway/internal/server"
	"api-gateway/pkg/logger"
)

func main() {
	cfg := config.GetConfig()

	l := logger.NewLogger(cfg.ToLoggerConfig())

	webApp := server.NewApp(cfg, l)
	webApp.Start()
}
