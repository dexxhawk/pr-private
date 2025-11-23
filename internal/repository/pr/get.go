package pr

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	model_pr "github.com/dexxhawk/pr-private/internal/repository/models/pr"
	schema_pr "github.com/dexxhawk/pr-private/internal/repository/schema/pr"
)

func (r *Repo) GetPRByID(ctx context.Context, prID string) (*model_pr.PR, error) {
	query := r.queryBuilder.
		Select(
			schema_pr.ColumnID,
			schema_pr.ColumnName,
			schema_pr.ColumnAuthorID,
			schema_pr.ColumnStatus,
			schema_pr.ColumnCreatedAt,
			schema_pr.ColumnMergedAt,
		).
		From(schema_pr.Table).
		Where(sq.Eq{schema_pr.ColumnID: prID})

	var modelPr model_pr.PR
	err := r.runner.Getx(ctx, &modelPr, query)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrPRNotFound
		}
		return nil, fmt.Errorf("get PR by ID: %w", err)
	}
	return &modelPr, nil
}
