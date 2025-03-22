package sublists

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type SublistsRepository struct {
	Dbpool       *pgxpool.Pool
	CusotmLogger *zerolog.Logger
}

func NewSublistsRepository(dbpool *pgxpool.Pool, customLogger *zerolog.Logger) *SublistsRepository {
	return &SublistsRepository{
		Dbpool:       dbpool,
		CusotmLogger: customLogger,
	}
}

func (r *SublistsRepository) GetAll() ([]Sublist, error) {
	query := "select * from sublists"
	rows, err := r.Dbpool.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	sublists, err := pgx.CollectRows(rows, pgx.RowToStructByName[Sublist])
	if err != nil {
		return nil, err
	}
	return sublists, nil
}
