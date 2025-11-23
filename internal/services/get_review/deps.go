package get_review

import (
	"context"

	model_pr "github.com/dexxhawk/pr-private/internal/repository/models/pr"
)

type ReviewerRepo interface {
	GetUserReviewPRs(ctx context.Context, userID string) ([]model_pr.PR, error)
}
