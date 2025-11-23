package reviewer

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"

	schema_reviewer "github.com/dexxhawk/pr-private/internal/repository/schema/reviewer"
)

func (r *Repo) ReplaceReviewer(ctx context.Context, prID string, oldReviewerID string, newReviewerID string) error {
	query := r.queryBuilder.
		Update(schema_reviewer.Table).
		Set(schema_reviewer.ColumnUserID, newReviewerID).
		Where(sq.And{
			sq.Eq{schema_reviewer.ColumnPRID: prID},
			sq.Eq{schema_reviewer.ColumnUserID: oldReviewerID},
		})

	_, err := r.runner.Execx(ctx, query)
	if err != nil {
		return fmt.Errorf("exec runner: %w", err)
	}
	return nil
}