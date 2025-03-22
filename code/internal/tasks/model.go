package tasks

import "time"

type TaskForm struct {
	ListId     int
	Title      string
	Desc       string
	Context    string
	Scenario   string
	Criterion  string
	StatusId   int
	Deadline   time.Time
	SublistId  int
	PriorityId int
}

type Task struct {
	Id         int       `db:"id"`
	ListId     int       `db:"list_id"`
	Title      string    `db:"title"`
	Desc       string    `db:"desc,omitempty"`
	Context    string    `db:"context,omitempty"`
	Scenario   string    `db:"scenario,omitempty"`
	Criterion  string    `db:"criterion,omitempty"`
	StatusId   int       `db:"status_id"`
	Deadline   time.Time `db:"deadline,omitempty"`
	SublistId  int       `db:"sublist_id"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at,omitempty"`
	PriorityId int       `db:"priority_id"`
	IsDel      int       `db:"is_del"`
}
