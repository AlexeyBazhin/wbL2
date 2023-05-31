package repository

import (
	"context"
	"fmt"
	"net/http"

	"dev11/internal/domain/model"
	"dev11/internal/errors"

	sq "github.com/Masterminds/squirrel"
)

func (r *Repository) InsertUser(ctx context.Context, user *model.User) *errors.MyErr {
	row := r.sq.Insert("users").
		Columns("email",
			"first_name", "last_name",
			"birth_date", "hashed_password").
		Values(user.Email,
			user.FirstName, user.LastName,
			user.BirthDate, user.HashedPassword).
		Suffix("RETURNING \"id\"").
		QueryRowContext(ctx)

	if err := row.Scan(&user.ID); err != nil {
		return &errors.MyErr{
			Code: http.StatusInternalServerError,
			Err:  fmt.Errorf("error while scanning sql row: %w", err),
		}
	}
	return nil
}

func (r *Repository) GetUser(ctx context.Context, id int) (*model.User, *errors.MyErr) {
	row := r.sq.Select("id",
		"email", "first_name", "last_name",
		"birth_date", "hashed_password").
		From("users u").
		Where(sq.Eq{"u.id": id}).
		QueryRowContext(ctx)

	user := &model.User{}
	if err := row.Scan(&user.ID, &user.Email,
		&user.FirstName, &user.LastName,
		&user.BirthDate, &user.HashedPassword); err != nil {
		return nil, &errors.MyErr{
			Code: http.StatusInternalServerError,
			Err:  fmt.Errorf("error while scanning sql row: %w", err),
		}
	}
	return user, nil
}
