package app

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/if-i-am-gone/if-i-am-gone-web-backend/internal/config"
	"github.com/if-i-am-gone/if-i-am-gone-web-backend/internal/database/pg"
	"github.com/if-i-am-gone/if-i-am-gone-web-backend/internal/services/auth/login"
)

func Run() error {
	errConfigLoad := config.Load()
	if errConfigLoad != nil {
		return fmt.Errorf("app error on Config.Load: %w", errConfigLoad)
	}

	cfg, errConfigGet := config.Get()
	if errConfigGet != nil {
		return fmt.Errorf("app error on Config.Get: %w", errConfigGet)
	}

	errPostgres := initializePostgres(cfg)
	if errPostgres != nil {
		return fmt.Errorf("app error on initializePostgres: %w", errPostgres)
	}

	runApiServer(cfg)

	return nil
}

func initializePostgres(cfg *config.Config) error {
	_, err := pg.ConnectPostgres(context.Background(), fmt.Sprintf("postgresql://%s:%d/%s?user=%s&password=%s",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.DB_Name,
		cfg.Postgres.User,
		cfg.Postgres.Password))

	if err != nil {
		return fmt.Errorf("error on initializing postgres: %w", err)
	}

	return nil
}

func runApiServer(cfg *config.Config) {
	r := gin.Default()

	r.GET("/auth/login", login.Login)

	r.Run(fmt.Sprintf("%s:%d",
		cfg.Api.Hostname,
		cfg.Api.Port))
}
