package get_review

import (
	"context"
	"fmt"

	"github.com/dexxhawk/pr-private/internal/domain"
	"github.com/dexxhawk/pr-private/internal/repository/models/pr"
)

func (s *Service) GetUserReviewPRs(ctx context.Context, userID string) ([]domain.PR, error) {

	prModels, err := s.reviewerRepo.GetUserReviewPRs(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("get user review PRs: %w", err)
	}
	prDomains := pr.PR{}.Domains(prModels)
	return prDomains, nil
}
