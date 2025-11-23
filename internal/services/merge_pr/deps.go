package merge_pr

import (
	"context"

	mpr "github.com/dexxhawk/pr-private/internal/repository/models/pr"
	mreviewer "github.com/dexxhawk/pr-private/internal/repository/models/reviewer"
)

type PRRepo interface {
	MergePR(ctx context.Context, prID string) (*mpr.PR, error)
}

type ReviewerRepo interface {
	GetUserByPR(ctx context.Context, prID string) ([]mreviewer.Reviewer, error)
}
