package database

import (
	"context"
	"go/context-todo/config"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

func CreateDbPool(config *config.DatabaseConfig, logger *zerolog.Logger) *pgxpool.Pool {
	dbpool, err := pgxpool.New(context.Background(), config.Url)
	if err != nil {
		logger.Error().Msg("DB connection failed")
		panic(err)
	}
	logger.Info().Msg("DB connected")
	return dbpool
}
