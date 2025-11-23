package http

import (
	"context"
	"errors"
	"fmt"

	"github.com/dexxhawk/pr-private/internal/domain"
	"github.com/dexxhawk/pr-private/internal/generated/api"
	"github.com/dexxhawk/pr-private/internal/services/api_mapper"
)

type MergePRSrv interface {
	MergePR(ctx context.Context, prID string) (*domain.PR, []string, error)
}

func (h *Server) PullRequestMergePost(ctx context.Context, req *api.PullRequestMergePostReq) (api.PullRequestMergePostRes, error) {

	pullReq, reviewerIDs, err := h.mergePRSrv.MergePR(ctx, req.PullRequestID)

	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return &api.ErrorResponse{
				Error: api.ErrorResponseError{
					Code:    api.ErrorResponseErrorCodeNOTFOUND,
					Message: "resource not found",
				},
			}, nil
		}
		return nil, fmt.Errorf("merge pr: %w", err)
	}

	apiPR := api_mapper.PRDomainToApi(*pullReq, reviewerIDs)

	resp := api.PullRequestMergePostOK{
		Pr: api.NewOptPullRequest(apiPR),
	}

	return &resp, nil
}
