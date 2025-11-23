package pr

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	model "github.com/dexxhawk/pr-private/internal/repository/models/pr"
	schema_pr "github.com/dexxhawk/pr-private/internal/repository/schema/pr"
)

const PRStatusMerged = 1

var ErrPRNotFound = errors.New("pull request not found")

func (r *Repo) MergePR(ctx context.Context, prID string) (*model.PR, error) {
	query := r.queryBuilder.
		Update(schema_pr.Table).
		Set(schema_pr.ColumnStatus, PRStatusMerged).
		// Обновляем merged_at, только если статус менялся
		Set(schema_pr.ColumnMergedAt,
			sq.Expr(`CASE WHEN status != ? THEN ? ELSE merged_at END`,
				PRStatusMerged, time.Now()),
		).
		Where(sq.Eq{schema_pr.ColumnID: prID}).
		Suffix("RETURNING " + schema_pr.ColumnID + ", " + schema_pr.ColumnName + ", " + schema_pr.ColumnAuthorID + ", " + schema_pr.ColumnStatus +
			", " + schema_pr.ColumnCreatedAt + ", " + schema_pr.ColumnMergedAt)

	var pullReq model.PR
	err := r.runner.Getx(ctx, &pullReq, query)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrPRNotFound
		}
		return nil, fmt.Errorf("exec runner: %w", err)
	}
	return &pullReq, nil
}
