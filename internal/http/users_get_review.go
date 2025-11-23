package http

import (
	"context"
	"fmt"

	"github.com/dexxhawk/pr-private/internal/domain"
	"github.com/dexxhawk/pr-private/internal/generated/api"
	"github.com/dexxhawk/pr-private/internal/services/api_mapper"
)

type GetUserReviewPRsSrv interface {
	GetUserReviewPRs(ctx context.Context, userID string) ([]domain.PR, error)
}

func (h *Server) UsersGetReviewGet(ctx context.Context, params api.UsersGetReviewGetParams) (*api.UsersGetReviewGetOK, error) {

	pullReqs, err := h.getUserReviewPRsSrv.GetUserReviewPRs(ctx, params.UserID)
	if err != nil {
		return nil, fmt.Errorf("get user review PRs: %w", err)
	}

	apiPullReqs := api_mapper.PRsDomainToPullRequestsShortApi(pullReqs)
	resp := api.UsersGetReviewGetOK{
		UserID:       params.UserID,
		PullRequests: apiPullReqs,
	}
	return &resp, nil
}
