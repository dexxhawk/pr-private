package merge_pr

import (
	"context"
	"errors"
	"fmt"

	"github.com/dexxhawk/pr-private/internal/domain"
	repo_pr "github.com/dexxhawk/pr-private/internal/repository/pr"
)

func (s *Service) MergePR(ctx context.Context, prID string) (*domain.PR, []string, error) {
	pullReq, err := s.prRepo.MergePR(ctx, prID)
	if err != nil {
		if errors.Is(err, repo_pr.ErrPRNotFound) {
			return nil, nil, domain.ErrNotFound
		}
		return nil, nil, fmt.Errorf("merge PR: %w", err)
	}
	users, err := s.reviewerRepo.GetUserByPR(ctx, pullReq.ID)
	if err != nil {
		return nil, nil, fmt.Errorf("get reviewers by pr_id: %w", err)
	}

	reviewersIDs := make([]string, 0, len(users))
	for _, user := range users {
		reviewersIDs = append(reviewersIDs, user.UserID)
	}

	domainPR := pullReq.Domain()
	return &domainPR, reviewersIDs, nil
}
