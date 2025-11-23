package user

import (
	"context"
	"fmt"

	model "github.com/dexxhawk/pr-private/internal/repository/models/user"
	schema "github.com/dexxhawk/pr-private/internal/repository/schema/user"
)

// если пользак существует (по id проверка), то обновит все переданные  ID
func (r *Repo) InsertOrUpdateUsers(ctx context.Context, users ...model.User) error {
	query := r.queryBuilder.
		Insert(schema.Table).
		Columns(
			schema.ColumnID,
			schema.ColumnName,
			schema.ColumnIsActive,
			schema.ColumnTeamName,
		)

	for _, user := range users {
		query = query.Values(
			user.ID,
			user.Name,
			user.IsActive,
			user.TeamName,
		)
	}

	suffix := `ON CONFLICT(id) DO UPDATE SET id = excluded.id, "name" = excluded."name", is_active = excluded.is_active, team_name = excluded.team_name`
	query = query.Suffix(suffix)

	_, err := r.runner.Execx(ctx, query)
	if err != nil {
		return fmt.Errorf("exec runner: %w", err)
	}
	return nil
}
