package priority

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type PriorityRepository struct {
	Dbpool       *pgxpool.Pool
	CusotmLogger *zerolog.Logger
}

func NewPriorityRepository(dbpool *pgxpool.Pool, customLogger *zerolog.Logger) *PriorityRepository {
	return &PriorityRepository{
		Dbpool:       dbpool,
		CusotmLogger: customLogger,
	}
}

func (r *PriorityRepository) GetAll() ([]Priority, error) {
	query := "select * from priority"
	rows, err := r.Dbpool.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	priorities, err := pgx.CollectRows(rows, pgx.RowToStructByName[Priority])
	if err != nil {
		return nil, err
	}
	return priorities, nil
}
