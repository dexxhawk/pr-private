package add_team

import (
	"context"
	"errors"
	"fmt"

	"github.com/dexxhawk/pr-private/internal/domain"
	mteam "github.com/dexxhawk/pr-private/internal/repository/models/team"
	muser "github.com/dexxhawk/pr-private/internal/repository/models/user"
	"github.com/dexxhawk/pr-private/pkg/query_runner"
)

var ErrTeamAlreadyExists = errors.New("team already exists")

func (s *Service) AddTeam(ctx context.Context, team domain.Team, users []domain.User) error {

	err := s.txManager.Do(ctx, func(ctx context.Context) error {
		err := s.teamRepo.InsertTeam(ctx, mteam.Team{}.Model(team))
		if err != nil {
			if errors.Is(err, query_runner.ErrUnique) {
				err = ErrTeamAlreadyExists
			}
			return fmt.Errorf("insert team: %w", err)
		}

		err = s.userRepo.InsertOrUpdateUsers(ctx, muser.User{}.Models(users...)...)
		if err != nil {
			return fmt.Errorf("insert or update user: %w", err)
		}
 
		return nil
	})
	if err != nil {
		return fmt.Errorf("insert team && members in tx: %w", err)
	}

	return nil
}
