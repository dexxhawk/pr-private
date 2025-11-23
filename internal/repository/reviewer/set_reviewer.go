package reviewer

import (
	"context"
	"fmt"

	schema_reviewer "github.com/dexxhawk/pr-private/internal/repository/schema/reviewer"
)

func (r *Repo) SetUserReviewPRs(ctx context.Context, prID string, userIDs []string) error {
	query := r.queryBuilder.
		Insert(schema_reviewer.Table).
		Columns(
			schema_reviewer.ColumnPRID,
			schema_reviewer.ColumnUserID,
		)

	for _, reviewerID := range userIDs {
		query = query.Values(prID, reviewerID)
	}

	_, err := r.runner.Execx(ctx, query)
	if err != nil {
		return fmt.Errorf("exec runner: %w", err)
	}
	return nil
}
