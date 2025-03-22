package tasks

import (
	"context"
	"fmt"
	"go/context-todo/internal/list"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type TasksRepository struct {
	Dbpool       *pgxpool.Pool
	CusotmLogger *zerolog.Logger
}

func NewTasksRepository(dbpool *pgxpool.Pool, customLogger *zerolog.Logger) *TasksRepository {
	return &TasksRepository{
		Dbpool:       dbpool,
		CusotmLogger: customLogger,
	}
}

func (r *TasksRepository) createTask(form TaskForm) error {
	query := `
insert into tasks (
	list_id, title, "desc", context,
	scenario, criterion, status_id, deadline,
	sublist_id, priority_id
) values (
	@list_id, @title, @desc, @context,
	@scenario, @criterion, @status_id, @deadline,
	@sublist_id, @priority_id
)`
	args := pgx.NamedArgs{
		"list_id":     form.ListId,
		"title":       form.Title,
		"desc":        form.Desc,
		"context":     form.Context,
		"scenario":    form.Scenario,
		"criterion":   form.Criterion,
		"status_id":   form.StatusId,
		"deadline":    form.Deadline,
		"sublist_id":  form.SublistId,
		"priority_id": form.PriorityId,
	}
	_, err := r.Dbpool.Exec(context.Background(), query, args)
	if err != nil {
		return fmt.Errorf("impossible to create new task: %w", err)
	}
	r.CusotmLogger.Info().Msg("end of createTask() in repo")
	return nil
}

func (r *TasksRepository) GetAllListsForUser(userId int) ([]list.List, error) {
	query := "select * from lists l where l.user_id = $1"
	rows, err := r.Dbpool.Query(context.Background(), query, userId)
	if err != nil {
		return nil, err
	}
	lists, err := pgx.CollectRows(rows, pgx.RowToStructByName[list.List])
	if err != nil {
		return nil, err
	}
	return lists, nil
}
