package main

import (
	"api-gateway/internal/config"
	"api-gateway/internal/server"
	"api-gateway/pkg/logger"
)

func main() {
	cfg := config.GetConfig()

	l := logger.NewLogger(cfg.ToLoggerConfig())

	// TODO: broker for async requests
	//_, err := broker.NewNATSClient(cfg.ToNATSConfig())
	//if err != nil {
	//	l.Fatal(err)
	//}

	webApp := server.NewApp(cfg, l)
	webApp.Start()
}
