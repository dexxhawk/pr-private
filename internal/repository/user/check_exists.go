package user

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	schema "github.com/dexxhawk/pr-private/internal/repository/schema/user"
)

func (r *Repo) CheckUserExists(ctx context.Context, userID string) (bool, error) {
	query := r.queryBuilder.
		Select("1").
		Prefix("SELECT EXISTS (").
		From(schema.Table).
		Where(sq.Eq{schema.ColumnID: userID}).
		Suffix(")")

	var exists bool
	err := r.runner.Getx(ctx, &exists, query)
	if err != nil {
		return false, fmt.Errorf("exec runner select: %w", err)
	}

	return exists, nil
}
