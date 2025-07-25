package main

import (
	"github.com/mauricioromagnollo/kafrest/config"
	"github.com/mauricioromagnollo/kafrest/pkg/logger"
)

func main() {
	// Build config instance
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	// Build logger instance
	log := logger.NewLogger(cfg.App.Name, cfg.App.Env, "info")

	// First log message
	log.Info("application started")

	// ...

	log.Info("kafrest api started", logger.ExtraProps{
		"port":         cfg.API.Port,
		"kafka_broker": cfg.Kafka.BrokerConnect,
	})
}
