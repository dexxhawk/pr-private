package get_team

import (
	"context"

	muser "github.com/dexxhawk/pr-private/internal/repository/models/user"
)

type TeamRepo interface {
	CheckTeamExists(ctx context.Context, teamName string) (bool, error)
}

type UserRepo interface {
	GetUsersByTeam(ctx context.Context, teamName string) ([]muser.User, error)
}
