package http

import (
	"context"
	"errors"
	"fmt"

	"github.com/dexxhawk/pr-private/internal/domain"
	"github.com/dexxhawk/pr-private/internal/generated/api"
	"github.com/dexxhawk/pr-private/internal/services/add_team"
	"github.com/dexxhawk/pr-private/internal/services/api_mapper"
)

type AddTeamSrv interface {
	AddTeam(ctx context.Context, team domain.Team, users []domain.User) error
}

func (h *Server) TeamAddPost(ctx context.Context, req *api.Team) (api.TeamAddPostRes, error) {

	team, users := api_mapper.TeamApiToDomain(*req)

	err := h.addTeamSrv.AddTeam(ctx, team, users)

	if err != nil {
		if errors.Is(err, add_team.ErrTeamAlreadyExists) {
			return &api.ErrorResponse{
				Error: api.ErrorResponseError{
					Code:    api.ErrorResponseErrorCodeTEAMEXISTS,
					Message: "team_name already exists",
				},
			}, nil
		}
		return nil, fmt.Errorf("add team: %w", err)
	}

	resp := api.TeamAddPostCreated{
		Team: api.OptTeam{
			Value: *req,
			Set:   true,
		},
	}

	return &resp, nil
}
