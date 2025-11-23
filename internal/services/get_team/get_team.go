package get_team

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/dexxhawk/pr-private/internal/domain"
	models_user "github.com/dexxhawk/pr-private/internal/repository/models/user"
)

var ErrTeamNotFound = errors.New("team already exists")

func (s *Service) GetTeam(ctx context.Context, teamName string) ([]domain.User, error) {

	exists, err := s.teamRepo.CheckTeamExists(ctx, teamName)
	if err != nil {
		return nil, fmt.Errorf("get team by name: %w", err)
	}
	if !exists {
		return nil, ErrTeamNotFound
	}

	users, err := s.userRepo.GetUsersByTeam(ctx, teamName)
	if err != nil {
		return nil, fmt.Errorf("get users by team: %w", err)
	}

	slog.Info("users after get:", slog.Any("users", users))

	
	return models_user.User{}.Domains(users), nil

}
