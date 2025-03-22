package statuses

type Status struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
}
