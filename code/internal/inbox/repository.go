package inbox

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type InboxRepository struct {
	Dbpool       *pgxpool.Pool
	CusotmLogger *zerolog.Logger
}

func NewInboxRepository(dbpool *pgxpool.Pool, customLogger *zerolog.Logger) *InboxRepository {
	return &InboxRepository{
		Dbpool:       dbpool,
		CusotmLogger: customLogger,
	}
}

func (r *InboxRepository) getAllInboxTasks(userId int) ([]InboxTask, error) {
	query := `select t.id, t.title from tasks t where t.list_id in (
	select l.id from lists l
	where l.user_id = $1
) and t.sublist_id = 0`
	rows, err := r.Dbpool.Query(context.Background(), query, userId)
	if err != nil {
		return nil, err
	}
	lists, err := pgx.CollectRows(rows, pgx.RowToStructByName[InboxTask])
	if err != nil {
		return nil, err
	}
	return lists, nil
}
