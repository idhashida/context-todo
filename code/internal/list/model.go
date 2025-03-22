package list

import "time"

type ListForm struct {
	Name  string
	Color string
}

type List struct {
	Id        int       `db:"id"`
	Name      string    `db:"name"`
	UserId    int       `db:"user_id"`
	Color     string    `db:"color"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at,omitempty"`
}
