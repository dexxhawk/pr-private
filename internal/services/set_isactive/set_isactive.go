package set_isactive

import (
	"context"
	"errors"
	"fmt"

	"github.com/dexxhawk/pr-private/internal/domain"
	repo_user "github.com/dexxhawk/pr-private/internal/repository/user"
)

var ErrTeamNotFound = errors.New("team already exists")

func (s *Service) SetIsActive(ctx context.Context, userID string, isActive bool) (*domain.User, error) {

	user, err := s.userRepo.SetIsActive(ctx, userID, isActive)
	if err != nil {
		if errors.Is(err, repo_user.ErrUserNotFound) {
			return nil, domain.ErrNotFound
		}
		return nil, fmt.Errorf("set is active: %w", err)
	}
	domainUser := user.Domain()
	return &domainUser, nil

}
