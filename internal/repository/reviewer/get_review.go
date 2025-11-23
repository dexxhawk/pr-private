package reviewer

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	model_pr "github.com/dexxhawk/pr-private/internal/repository/models/pr"
	schema_pr "github.com/dexxhawk/pr-private/internal/repository/schema/pr"
	schema_reviewer "github.com/dexxhawk/pr-private/internal/repository/schema/reviewer"
)

func (r *Repo) GetUserReviewPRs(ctx context.Context, userID string) ([]model_pr.PR, error) {
	joinCondition := schema_pr.Table + "." + schema_pr.ColumnID + " = " + 
		schema_reviewer.Table + "." + schema_reviewer.ColumnPRID
	query := r.queryBuilder.
		Select(
			schema_pr.ColumnID,
			schema_pr.ColumnName,
			schema_pr.ColumnAuthorID,
			schema_pr.ColumnStatus,
		).
		From(schema_pr.Table).
		Join(schema_reviewer.Table + " ON " + joinCondition).
		Where(sq.Eq{schema_reviewer.Table + "." + schema_reviewer.ColumnUserID: userID})

	var modelPr []model_pr.PR
	err := r.runner.Selectx(ctx, &modelPr, query)
	if err != nil {
		return nil, fmt.Errorf("exec runner: %w", err)
	}
	return modelPr, nil
}
