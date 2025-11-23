package create_pr

import (
	"context"

	mpr "github.com/dexxhawk/pr-private/internal/repository/models/pr"
	muser "github.com/dexxhawk/pr-private/internal/repository/models/user"
)

type TeamRepo interface {
	CheckTeamExists(ctx context.Context, teamName string) (bool, error)
}

type UserRepo interface {
	CheckUserExists(ctx context.Context, userID string) (bool, error)
	GetUsersByTeam(ctx context.Context, teamName string) ([]muser.User, error)
	GetUserByID(ctx context.Context, userID string) (*muser.User, error)
}
type PRRepo interface {
	CreatePR(ctx context.Context, prID string, prName string, userID string) (*mpr.PR, error)
}
type ReviewerRepo interface {
	SetUserReviewPRs(ctx context.Context, prID string, userIDs []string) error
}
