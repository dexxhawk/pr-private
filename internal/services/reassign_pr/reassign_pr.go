package reassign_pr

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/dexxhawk/pr-private/internal/domain"
	mreviewer "github.com/dexxhawk/pr-private/internal/repository/models/reviewer"
	muser "github.com/dexxhawk/pr-private/internal/repository/models/user"
	repo_pr "github.com/dexxhawk/pr-private/internal/repository/pr"
)

const (
	MaxReviewersCount = 2
)

var (
	ErrPRAlreadyMerged = errors.New("PR is already merged")
	ErrNotAssign       = errors.New("this user was not assigned to this PR")
	ErrNoCandidate     = errors.New("no candidates available for reassignment")
)

func (s *Service) ReassignPR(ctx context.Context, prID string, oldUserID string) (*domain.PR, []string, *string, error) {
	pullReq, err := s.prRepo.GetPRByID(ctx, prID)
	if err != nil {
		if errors.Is(err, repo_pr.ErrPRNotFound) {
			return nil, nil, nil, domain.ErrNotFound
		}
		return nil, nil, nil, fmt.Errorf("get PR by ID: %w", err)
	}
	if pullReq.Status == repo_pr.PRStatusMerged {
		return nil, nil, nil, ErrPRAlreadyMerged
	}

	reviewers, err := s.reviewerRepo.GetUserByPR(ctx, pullReq.ID)

	found := false
	for _, reviewer := range reviewers {
		if reviewer.UserID == oldUserID {
			found = true
			break
		}
	}
	if !found {
		return nil, nil, nil, ErrNotAssign
	}

	user, err := s.userRepo.GetUserByID(ctx, oldUserID)
	if err != nil {
		if errors.Is(err, repo_pr.ErrPRNotFound) {
			return nil, nil, nil, domain.ErrNotFound
		}
	}
	teamUsers, err := s.userRepo.GetUsersByTeam(ctx, user.TeamName)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get users by team: %w", err)
	}

	filteredUsers := FilterUsers(teamUsers, pullReq.AuthorID, oldUserID)
	availableReviewers := SelectRandomReviewers(filteredUsers, MaxReviewersCount)
	reviewersIDs := ExtractIDsFromUsers(availableReviewers)

	if len(reviewersIDs) == 0 {
		return nil, nil, nil, ErrNoCandidate
	}

	reviewerID := reviewersIDs[0]
	err = s.reviewerRepo.ReplaceReviewer(ctx, prID, oldUserID, reviewerID)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("replace reviewer: %w", err)
	}

	newReviewers, err := s.reviewerRepo.GetUserByPR(ctx, prID)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get updated PR by ID: %w", err)
	}

	domainPR := pullReq.Domain()
	reviewerIDs := ExtractIDsFromReviewers(newReviewers)

	return &domainPR, reviewerIDs, &reviewerID, nil

}

func FilterUsers(users []muser.User, authorID string, oldUserID string) []muser.User {
	filteredUsers := make([]muser.User, 0, len(users))

	for _, user := range users {
		if user.ID == authorID || !user.IsActive || user.ID == oldUserID {
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

func ExtractIDsFromUsers(users []muser.User) []string {
	idS := make([]string, 0, len(users))
	for _, user := range users {
		idS = append(idS, user.ID)
	}
	return idS
}

func ExtractIDsFromReviewers(reviewers []mreviewer.Reviewer) []string {
	idS := make([]string, 0, len(reviewers))
	for _, reviewer := range reviewers {
		idS = append(idS, reviewer.UserID)
	}
	return idS
}
