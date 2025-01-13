package config

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/arafetki/go-echo-boilerplate/internal/env"
)

type Config struct {
	// Application configuration
	App struct {
		Name        string
		Version     string
		Env         string
		Description string
		Debug       bool
	}
	// Server configuration
	Server struct {
		Addr          string
		ReadTimeout   time.Duration
		WriteTimeout  time.Duration
		ShutdowPeriod time.Duration
	}
	// Logger configuration
	Logger struct {
		Level slog.Level
	}
	// Database configuration
	Database struct {
		Dsn         string
		Automigrate bool
	}
	// JWT configuration
	JWT struct {
		Key string
	}
}

func Load() Config {
	var cfg Config

	cfg.App.Name = "Go Echo Boilerplate"
	cfg.App.Version = "0.1.0"
	cfg.App.Env = env.GetString("APP_ENV", "development")
	cfg.App.Description = "A boilerplate for building RESTful APIs with Go and Echo"
	cfg.App.Debug = true
	cfg.Server.Addr = fmt.Sprintf(":%d", env.GetInt("SERVER_PORT", 8080))
	cfg.Server.ReadTimeout = env.GetDuration("SERVER_READ_TIMEOUT", 10*time.Second)
	cfg.Server.WriteTimeout = env.GetDuration("SERVER_WRITE_TIMEOUT", 30*time.Second)
	cfg.Server.ShutdowPeriod = env.GetDuration("SERVER_SHUTDOWN_PERIOD", 60*time.Second)
	cfg.Database.Dsn = env.GetString("DATABASE_DSN", "postgres:postgres@localhost:5432/postgres?sslmode=disable")
	cfg.Database.Automigrate = true
	cfg.Logger.Level = slog.LevelDebug
	if cfg.App.Env == "production" {
		cfg.App.Debug = false
		cfg.Logger.Level = slog.LevelInfo
		cfg.Database.Automigrate = false
	}
	cfg.JWT.Key = os.Getenv("JWT_KEY")

	return cfg
}
