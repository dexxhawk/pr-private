package api_mapper

import (
	"github.com/dexxhawk/pr-private/internal/domain"
	"github.com/dexxhawk/pr-private/internal/generated/api"
)

func TeamApiToDomain(team api.Team) (domain.Team, []domain.User) {
	domainTeam := domain.Team{
		Name: team.TeamName,
	}
	domainUsers := make([]domain.User, 0, len(team.Members))
	for _, user := range team.Members {
		domainUser := UserApiToDomain(user, team.TeamName)
		domainUsers = append(domainUsers, domainUser)
	}
	return domainTeam, domainUsers
}

func UserApiToDomain(user api.TeamMember, teamName string) domain.User {
	domainUser := domain.User{
		ID:       user.UserID,
		Name:     user.Username,
		IsActive: user.IsActive,
		TeamName: teamName,
	}
	return domainUser
}

func UserDomainToTeamMemberApi(user domain.User) api.TeamMember {
	apiTeamMember := api.TeamMember{
		UserID:   user.ID,
		Username: user.Name,
		IsActive: user.IsActive,
	}
	return apiTeamMember
}

func UserDomainToApi(user domain.User) api.User {
	apiUser := api.User{
		UserID:   user.ID,
		Username: user.Name,
		TeamName: user.TeamName,
		IsActive: user.IsActive,
	}
	return apiUser
}

func TeamDomainToApi(teamName string, users []domain.User) api.Team {
	apiUsers := make([]api.TeamMember, 0, len(users))
	for _, user := range users {
		apiUser := UserDomainToTeamMemberApi(user)
		apiUsers = append(apiUsers, apiUser)
	}

	apiTeam := api.Team{
		TeamName: teamName,
		Members:  apiUsers,
	}
	return apiTeam
}

func PRDomainToPullRequestShortApi(pr domain.PR) api.PullRequestShort {
	status := api.PullRequestShortStatusOPEN
	switch pr.Status {
	case 0:
		status = api.PullRequestShortStatusOPEN
	case 1:
		status = api.PullRequestShortStatusMERGED
	}

	apiPR := api.PullRequestShort{
		PullRequestID:   pr.ID,
		PullRequestName: pr.Name,
		AuthorID:        pr.AuthorID,
		Status:          status,
	}

	return apiPR
}

func PRsDomainToPullRequestsShortApi(pullReqs []domain.PR) []api.PullRequestShort {
	apiPullReqs := make([]api.PullRequestShort, 0, len(pullReqs))
	for _, pr := range pullReqs {
		apiPR := PRDomainToPullRequestShortApi(pr)
		apiPullReqs = append(apiPullReqs, apiPR)
	}
	return apiPullReqs
}

func PRDomainToApi(pr domain.PR, reviewersIDs []string) api.PullRequest {
	status := api.PullRequestStatusOPEN
	switch pr.Status {
	case 0:
		status = api.PullRequestStatusOPEN
	case 1:
		status = api.PullRequestStatusMERGED
	}

	var mergedAt api.OptNilDateTime
	if pr.MergedAt != nil {
		mergedAt = api.NewOptNilDateTime(*pr.MergedAt)
	}

	apiPR := api.PullRequest{
		PullRequestID:     pr.ID,
		PullRequestName:   pr.Name,
		AuthorID:          pr.AuthorID,
		Status:            status,
		AssignedReviewers: reviewersIDs,
		CreatedAt:         api.NewOptNilDateTime(pr.CreatedAt),
		MergedAt:          mergedAt,
	}

	return apiPR
}
