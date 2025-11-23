package add_team

import (
	"context"

	mteam "github.com/dexxhawk/pr-private/internal/repository/models/team"
	muser "github.com/dexxhawk/pr-private/internal/repository/models/user"
)

type TeamRepo interface {
	InsertTeam(ctx context.Context, team mteam.Team) error
}

type UserRepo interface {
	InsertOrUpdateUsers(ctx context.Context, users ...muser.User) error
}
