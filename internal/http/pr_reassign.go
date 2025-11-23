package http

import (
	"context"
	"errors"
	"fmt"

	"github.com/dexxhawk/pr-private/internal/domain"
	"github.com/dexxhawk/pr-private/internal/generated/api"
	"github.com/dexxhawk/pr-private/internal/services/api_mapper"
	"github.com/dexxhawk/pr-private/internal/services/reassign_pr"
)

type ReassignPRSrv interface {
	ReassignPR(ctx context.Context, prID string, oldUserID string) (*domain.PR, []string, *string, error)
}

func (h *Server) PullRequestReassignPost(ctx context.Context, req *api.PullRequestReassignPostReq) (api.PullRequestReassignPostRes, error) {

	pullReq, reviewersIDs, newReviewerID, err := h.reassignPRSrv.ReassignPR(ctx, req.PullRequestID, req.OldUserID)

	if err != nil {
		switch {
		case errors.Is(err, domain.ErrNotFound):
			return &api.PullRequestReassignPostNotFound{
				Error: api.ErrorResponseError{
					Code:    api.ErrorResponseErrorCodeNOTFOUND,
					Message: "resource not found",
				},
			}, nil
		case errors.Is(err, reassign_pr.ErrPRAlreadyMerged):
			return &api.PullRequestReassignPostConflict{
				Error: api.ErrorResponseError{
					Code:    api.ErrorResponseErrorCodePRMERGED,
					Message: "cannot reassign on merged PR",
				},
			}, nil
		case errors.Is(err, reassign_pr.ErrNotAssign):
			return &api.PullRequestReassignPostConflict{
				Error: api.ErrorResponseError{
					Code:    api.ErrorResponseErrorCodeNOTASSIGNED,
					Message: "reviewer is not assigned to this PR",
				},
			}, nil
		case errors.Is(err, reassign_pr.ErrNoCandidate):
			return &api.PullRequestReassignPostConflict{
				Error: api.ErrorResponseError{
					Code:    api.ErrorResponseErrorCodeNOCANDIDATE,
					Message: "no active replacement candidate in team",
				},
			}, nil
		default:
			return nil, fmt.Errorf("reassign pr: %w", err)
		}
	}

	apiPR := api_mapper.PRDomainToApi(*pullReq, reviewersIDs)

	return &api.PullRequestReassignPostOK{
		Pr:         apiPR,
		ReplacedBy: *newReviewerID,
	}, nil

}
