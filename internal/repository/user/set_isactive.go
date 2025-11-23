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

func (r *Repo) SetIsActive(ctx context.Context, userID string, isActive bool) (*model.User, error) {
	query := r.queryBuilder.
		Update(schema.Table).
		Set(schema.ColumnIsActive, isActive).
		Where(sq.Eq{schema.ColumnID: userID}).
		Suffix("RETURNING " + schema.ColumnID + ", " + schema.ColumnName + ", " + schema.ColumnIsActive + ", " + schema.ColumnTeamName)

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
