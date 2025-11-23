package user

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	model "github.com/dexxhawk/pr-private/internal/repository/models/user"
	schema "github.com/dexxhawk/pr-private/internal/repository/schema/user"
)

func (r *Repo) GetUsersByTeam(ctx context.Context, teamName string) ([]model.User, error) {
	query := r.queryBuilder.
		Select(
			schema.ColumnID,
			schema.ColumnName,
			schema.ColumnIsActive,
			schema.ColumnTeamName).
		From(schema.Table).
		Where(
			sq.Eq{
				schema.ColumnTeamName: teamName,
			},
		)

	var users []model.User
	err := r.runner.Selectx(ctx, &users, query)
	if err != nil {
		return nil, fmt.Errorf("exec runner: %w", err)
	}

	return users, nil
}
