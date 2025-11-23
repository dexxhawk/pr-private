package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	model "github.com/dexxhawk/pr-private/internal/repository/models/user"
	schema "github.com/dexxhawk/pr-private/internal/repository/schema/user"
)

func (r *Repo) GetUserByID(ctx context.Context, userID string) (*model.User, error) {
	query := r.queryBuilder.
		Select(
			schema.ColumnID,
			schema.ColumnName,
			schema.ColumnIsActive,
			schema.ColumnTeamName,
		).
		From(schema.Table).
		Where(sq.Eq{schema.ColumnID: userID})

	var user model.User
	err := r.runner.Getx(ctx, &user, query)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, fmt.Errorf("exec runner: %w", err)
	}
	return &user, nil
}
