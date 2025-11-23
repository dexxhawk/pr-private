package http

import (
	"context"
	"errors"
	"fmt"

	"github.com/dexxhawk/pr-private/internal/domain"
	"github.com/dexxhawk/pr-private/internal/generated/api"
	"github.com/dexxhawk/pr-private/internal/services/api_mapper"
	"github.com/dexxhawk/pr-private/internal/services/get_team"
)

type GetTeamSrv interface {
	GetTeam(ctx context.Context, teamName string) ([]domain.User, error)
}

func (h *Server) TeamGetGet(ctx context.Context, params api.TeamGetGetParams) (api.TeamGetGetRes, error) {

	teamName := params.TeamName

	users, err := h.getTeamSrv.GetTeam(ctx, teamName)

	if err != nil {
		if errors.Is(err, get_team.ErrTeamNotFound) {
			return &api.ErrorResponse{
				Error: api.ErrorResponseError{
					Code:    api.ErrorResponseErrorCodeNOTFOUND,
					Message: "resource not found",
				},
			}, nil
		}
		return nil, fmt.Errorf("get team: %w", err)
	}

	resp := api_mapper.TeamDomainToApi(teamName, users)

	return &resp, nil
}
