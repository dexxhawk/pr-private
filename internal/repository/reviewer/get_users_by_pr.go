package reviewer

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	mreviewer "github.com/dexxhawk/pr-private/internal/repository/models/reviewer"
	schema_reviewer "github.com/dexxhawk/pr-private/internal/repository/schema/reviewer"
)

func (r *Repo) GetUserByPR(ctx context.Context, prID string) ([]mreviewer.Reviewer, error) {
	query := r.queryBuilder.
		Select(
			schema_reviewer.ColumnPRID,
			schema_reviewer.ColumnUserID,
		).
		From(schema_reviewer.Table).
		Where(sq.Eq{schema_reviewer.ColumnPRID: prID})

	var users []mreviewer.Reviewer
	err := r.runner.Selectx(ctx, &users, query)
	if err != nil {
		return nil, fmt.Errorf("get reviewers: %w", err)
	}
	return users, nil
}
