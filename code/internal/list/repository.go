package list

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type ListRepository struct {
	Dbpool       *pgxpool.Pool
	CusotmLogger *zerolog.Logger
}

func NewListRepository(dbpool *pgxpool.Pool, customLogger *zerolog.Logger) *ListRepository {
	return &ListRepository{
		Dbpool:       dbpool,
		CusotmLogger: customLogger,
	}
}

func (r *ListRepository) createList(form ListForm, userId int) error {
	query := `insert into lists (name, color, user_id)
values (@name, @color, @user_id)`
	args := pgx.NamedArgs{
		"name":    form.Name,
		"color":   form.Color,
		"user_id": userId,
	}
	_, err := r.Dbpool.Exec(context.Background(), query, args)
	if err != nil {
		return fmt.Errorf("can't create new list: %w", err)
	}
	return nil
}
