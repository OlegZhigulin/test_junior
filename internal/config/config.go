package config

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type ServiceInformation struct {
	About           string `env:"ABOUT" env-default:""`
	Description     string `env:"DESCRIPTION" env-default:""`
	Stage           string `env:"STAGE" env-default:"develop"`
	StartUpDatetime time.Time
	Version         string `env:"VERSION" env-default:"1.0.0"`
}

type Config struct {
	HTTPServer struct {
		Host    string `env:"HTTP_SERVER_HOST" env-default:"localhost"`
		Port    string `env:"HTTP_SERVER_PORT" env-default:"8000"`
		Timeout string `env:"HTTP_SERVER_TIMEOUT" env-default:"10s"`
	}
	ServiceInformation ServiceInformation
	Logger             *slog.Logger
}

func NewServiceInformation() *ServiceInformation {
	return &ServiceInformation{
		About:           getEnvWithDefault("ABOUT", ""),
		Description:     getEnvWithDefault("DESCRIPTION", ""),
		Stage:           getEnvWithDefault("STAGE", "develop"),
		Version:         getEnvWithDefault("VERSION", "1.0.0"),
		StartUpDatetime: time.Now(),
	}
}

func GetConfigTimeoutServer(cfg *Config) time.Duration {
	duration, err := time.ParseDuration(cfg.HTTPServer.Timeout)
	if err != nil {
		return 10 * time.Second
	}
	return duration
}

func getEnvWithDefault(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func MustLoad() (*Config, error) {
	var cfg Config

	if err := godotenv.Load(); err != nil {
		log.Fatal("Warning: .env file not found or cannot be loaded: ", "error", err)
	}

	cfg.HTTPServer.Host = getEnvWithDefault("HTTP_SERVER_HOST", "localhost")
	cfg.HTTPServer.Port = getEnvWithDefault("HTTP_SERVER_PORT", "8000")
	cfg.HTTPServer.Timeout = getEnvWithDefault("HTTP_SERVER_TIMEOUT", "10s")

	serviceInfo := NewServiceInformation()
	cfg.ServiceInformation = *serviceInfo

	cfg.Logger = initLogger(cfg.ServiceInformation.Stage)

	cfg.Logger.Debug("Loaded config: ", slog.String("cfg", fmt.Sprintf("%+v", cfg)))

	return &cfg, nil
}
