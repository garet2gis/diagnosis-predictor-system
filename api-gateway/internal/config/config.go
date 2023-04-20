package config

import (
	"api-gateway/pkg/logger"
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPConfig struct {
	Port                           string `env:"PORT"  env-required:"true"`
	Host                           string `env:"HOST"  env-required:"true"`
	ReadTimeoutSeconds             int    `env:"READ_TIMEOUT_SEC" env-default:"0"`
	WriteTimeoutSeconds            int    `env:"WRITE_TIMEOUT_SEC" env-default:"0"`
	GracefulShutdownTimeoutSeconds int    `env:"GRACEFUL_SHUTDOWN_TIMEOUT_SEC" env-default:"15"`
}

type LoggerConfig struct {
	ReportCaller    bool   `env:"LOGGER_REPORT_CALLER" env-default:"true"`
	IsJSONFormatter bool   `env:"LOGGER_IS_JSON_FORMATTER" env-default:"false"`
	LoggingLevel    string `env:"LOGGING_LEVEL" env-default:"trace"`
}

type Config struct {
	HTTPConfig
	LoggerConfig
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		log.Print("Read application configuration")

		instance = &Config{}
		if err := cleanenv.ReadEnv(instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)

			log.Print(help)
			log.Fatal(err)
		}
	})

	return instance
}

func (c Config) ToLoggerConfig() logger.LoggerConfig {
	return logger.LoggerConfig{
		ReportCaller:    c.ReportCaller,
		IsJSONFormatter: c.IsJSONFormatter,
		LoggingLevel:    c.LoggingLevel,
	}
}
