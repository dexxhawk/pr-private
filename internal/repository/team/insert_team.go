package team

import (
	"context"
	"fmt"

	model "github.com/dexxhawk/pr-private/internal/repository/models/team"
	schema "github.com/dexxhawk/pr-private/internal/repository/schema/team"
)

func (r *Repo) InsertTeam(ctx context.Context, team model.Team) error {
	query := r.queryBuilder.
		Insert(schema.Table).
		Columns(schema.ColumnName).
		Values(team.Name)

	_, err := r.runner.Execx(ctx, query)
	if err != nil {
		return fmt.Errorf("exec runner: %w", err)
	}
	return nil
}
