package statuses

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type StatusesRepository struct {
	Dbpool       *pgxpool.Pool
	CusotmLogger *zerolog.Logger
}

func NewStatusesRepository(dbpool *pgxpool.Pool, customLogger *zerolog.Logger) *StatusesRepository {
	return &StatusesRepository{
		Dbpool:       dbpool,
		CusotmLogger: customLogger,
	}
}

func (r *StatusesRepository) GetAll() ([]Status, error) {
	query := "select * from statuses"
	rows, err := r.Dbpool.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	statuses, err := pgx.CollectRows(rows, pgx.RowToStructByName[Status])
	if err != nil {
		return nil, err
	}
	return statuses, nil
}
