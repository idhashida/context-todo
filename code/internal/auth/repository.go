package auth

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type AuthRepository struct {
	Dbpool       *pgxpool.Pool
	CusotmLogger *zerolog.Logger
}

func NewAuthRepository(dbpool *pgxpool.Pool, customLogger *zerolog.Logger) *AuthRepository {
	return &AuthRepository{
		Dbpool:       dbpool,
		CusotmLogger: customLogger,
	}
}

func (r *AuthRepository) createUser(form RegisterForm) error {
	query := `insert into users (username, email, password_hash)
values (@username, @email, @password_hash)`
	args := pgx.NamedArgs{
		"username":      form.Username,
		"email":         form.Email,
		"password_hash": form.Password,
	}
	_, err := r.Dbpool.Exec(context.Background(), query, args)
	if err != nil {
		return fmt.Errorf("registration is not possible: %w", err)
	}
	return nil
}
