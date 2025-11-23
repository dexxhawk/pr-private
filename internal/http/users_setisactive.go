package http

import (
	"context"
	"errors"
	"fmt"

	"github.com/dexxhawk/pr-private/internal/domain"
	"github.com/dexxhawk/pr-private/internal/generated/api"
	"github.com/dexxhawk/pr-private/internal/services/api_mapper"
)

type SetIsActiveSrv interface {
	SetIsActive(ctx context.Context, userID string, isActive bool) (*domain.User, error)
}

func (h *Server) UsersSetIsActivePost(ctx context.Context, req *api.UsersSetIsActivePostReq) (api.UsersSetIsActivePostRes, error) {

	user, err := h.setIsActiveSrv.SetIsActive(ctx, req.UserID, req.IsActive)

	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return &api.ErrorResponse{
				Error: api.ErrorResponseError{
					Code:    api.ErrorResponseErrorCodeNOTFOUND,
					Message: "resource not found",
				},
			}, nil
		}
		return nil, fmt.Errorf("set is_active: %w", err)
	}

	apiUser := api_mapper.UserDomainToApi(*user)
	resp := api.UsersSetIsActivePostOK{
		User: api.OptUser{
			Value: apiUser,
			Set: true,
		},
	}
	return &resp, nil
}
