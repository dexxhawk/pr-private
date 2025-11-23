package pr

import (
	"context"
	"fmt"
	"time"

	model_pr "github.com/dexxhawk/pr-private/internal/repository/models/pr"
	schema_pr "github.com/dexxhawk/pr-private/internal/repository/schema/pr"
)

func (r *Repo) CreatePR(ctx context.Context, prID string, prName string, userID string) (*model_pr.PR, error) {
	query := r.queryBuilder.
		Insert(schema_pr.Table).
		Columns(
			schema_pr.ColumnID,
			schema_pr.ColumnName,
			schema_pr.ColumnAuthorID,
			schema_pr.ColumnStatus,
			schema_pr.ColumnCreatedAt,
			schema_pr.ColumnMergedAt,
		).
		Values(prID, prName, userID, 0, time.Now(), nil).
		Suffix("RETURNING " + schema_pr.ColumnID + ", " + schema_pr.ColumnName + ", " + schema_pr.ColumnAuthorID + ", " + schema_pr.ColumnStatus +
			", " + schema_pr.ColumnCreatedAt + ", " + schema_pr.ColumnMergedAt)

	var modelPr model_pr.PR
	err := r.runner.Getx(ctx, &modelPr, query)

	if err != nil {
		return nil, fmt.Errorf("exec runner: %w", err)
	}
	return &modelPr, nil
}
