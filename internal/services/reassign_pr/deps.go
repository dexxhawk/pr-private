package reassign_pr

import (
	"context"

	mpr "github.com/dexxhawk/pr-private/internal/repository/models/pr"
	mreviewer "github.com/dexxhawk/pr-private/internal/repository/models/reviewer"
	muser "github.com/dexxhawk/pr-private/internal/repository/models/user"
)

type PRRepo interface {
	GetPRByID(ctx context.Context, prID string) (*mpr.PR, error)
}

type ReviewerRepo interface {
	GetUserByPR(ctx context.Context, prID string) ([]mreviewer.Reviewer, error)
	ReplaceReviewer(ctx context.Context, prID string, oldReviewerID string, newReviewerID string) error
}

type UserRepo interface {
	GetUsersByTeam(ctx context.Context, teamName string) ([]muser.User, error)
	GetUserByID(ctx context.Context, userID string) (*muser.User, error)
}
