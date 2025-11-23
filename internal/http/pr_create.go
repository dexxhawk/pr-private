package http

import (
	"context"
	"errors"
	"fmt"

	"github.com/dexxhawk/pr-private/internal/domain"
	"github.com/dexxhawk/pr-private/internal/generated/api"
	"github.com/dexxhawk/pr-private/internal/services/api_mapper"
	"github.com/dexxhawk/pr-private/internal/services/create_pr"
)

type CreatePRSrv interface {
	CreatePR(ctx context.Context, prID string, prName string, userID string) (*domain.PR, []string, error)
}

func (h *Server) PullRequestCreatePost(ctx context.Context, req *api.PullRequestCreatePostReq) (api.PullRequestCreatePostRes, error) {

	pullReq, reviewersIDs, err := h.createPRSrv.CreatePR(ctx, req.PullRequestID, req.PullRequestName, req.AuthorID)

	if err != nil {
		if errors.Is(err, create_pr.ErrUserNotFound) || errors.Is(err, create_pr.ErrTeamNotFound) {
			return &api.PullRequestCreatePostNotFound{
				Error: api.ErrorResponseError{
					Code:    api.ErrorResponseErrorCodeNOTFOUND,
					Message: "resource not found",
				},
			}, nil
		} else if errors.Is(err, create_pr.ErrPRAlrearyExists) {
			return &api.PullRequestCreatePostConflict{
				Error: api.ErrorResponseError{
					Code:    api.ErrorResponseErrorCodePREXISTS,
					Message: "PR id already exists",
				},
			}, nil
		}
		return nil, fmt.Errorf("create pr: %w", err)
	}

	apiPR := api_mapper.PRDomainToApi(*pullReq, reviewersIDs)

	resp := api.PullRequestCreatePostCreated{
		Pr: api.NewOptPullRequest(apiPR),
	}

	return &resp, nil
}
