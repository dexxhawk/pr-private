package create_pr

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/dexxhawk/pr-private/internal/domain"
	mpr "github.com/dexxhawk/pr-private/internal/repository/models/pr"
	muser "github.com/dexxhawk/pr-private/internal/repository/models/user"
	ruser "github.com/dexxhawk/pr-private/internal/repository/user"
	"github.com/dexxhawk/pr-private/pkg/query_runner"
)

const (
	MaxReviewersCount = 2
)

var (
	ErrUserNotFound    = errors.New("user not found")
	ErrTeamNotFound    = errors.New("team not found")
	ErrPRAlrearyExists = errors.New("pr already exists")
)

func FilterUsers(users []muser.User, authorID string) []muser.User {
	filteredUsers := make([]muser.User, 0, len(users))

	for _, user := range users {
		if user.ID == authorID || !user.IsActive {
			continue
		}
		filteredUsers = append(filteredUsers, user)
	}

	return filteredUsers
}

func SelectRandomReviewers(users []muser.User, maxCount int) []muser.User {
	if len(users) <= maxCount {
		return users
	}

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	indices := rnd.Perm(len(users))

	result := make([]muser.User, 0, maxCount)
	for i := range maxCount {
		result = append(result, users[indices[i]])
	}

	return result
}

func SelectRandomReviewersIDs(users []muser.User) []string {
	idS := make([]string, 0, len(users))
	for _, user := range users {
		idS = append(idS, user.ID)
	}
	return idS
}

func (s *Service) CreatePR(ctx context.Context, prID string, prName string, userID string) (*domain.PR, []string, error) {

	user, err := s.userRepo.GetUserByID(ctx, userID)
	if err != nil {
		if errors.Is(err, ruser.ErrUserNotFound) {
			return nil, nil, ErrUserNotFound
		}
		return nil, nil, fmt.Errorf("check user exists: %w", err)
	}

	exists, err := s.teamRepo.CheckTeamExists(ctx, user.TeamName)
	if err != nil {
		return nil, nil, fmt.Errorf("check team exists: %w", err)
	}

	if !exists {
		return nil, nil, ErrTeamNotFound
	}

	users, err := s.userRepo.GetUsersByTeam(ctx, user.TeamName)
	if err != nil {
		return nil, nil, fmt.Errorf("get users by team: %w", err)
	}
	// if len(users) == 0 {
	// 	return nil, ErrTeamUsersNotFound
	// } формально у нас может быть команда с 0 участниками

	filteredUsers := FilterUsers(users, userID)
	reviewers := SelectRandomReviewers(filteredUsers, MaxReviewersCount)
	reviewersIDs := SelectRandomReviewersIDs(reviewers)

	var pullReq *mpr.PR

	err = s.txManager.Do(ctx, func(ctx context.Context) error {
		pullReq, err = s.prRepo.CreatePR(ctx, prID, prName, userID)
		if err != nil {
			if errors.Is(err, query_runner.ErrUnique) {
				return ErrPRAlrearyExists
			}
			return err
		}
		if len(reviewersIDs) > 0 {
			s.reviewerRepo.SetUserReviewPRs(ctx, prID, reviewersIDs)
		}
		return nil
	})

	if err != nil {
		return nil, nil, fmt.Errorf("tx create pr: %w", err)
	}

	prDomain := pullReq.Domain()

	return &prDomain, reviewersIDs, nil
}
