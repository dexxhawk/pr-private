package http

import (
	"github.com/dexxhawk/pr-private/internal/generated/api"
)

type Server struct {
	addTeamSrv AddTeamSrv
	getTeamSrv GetTeamSrv
	setIsActiveSrv SetIsActiveSrv
	getUserReviewPRsSrv GetUserReviewPRsSrv
	createPRSrv CreatePRSrv
	mergePRSrv MergePRSrv
	reassignPRSrv ReassignPRSrv


	api.UnimplementedHandler
}

func New(
	addTeamSrv AddTeamSrv,
	getTeamSrv GetTeamSrv,
	setIsActiveSrv SetIsActiveSrv,
	getUserReviewPRsSrv GetUserReviewPRsSrv,
	createPRSrv CreatePRSrv,
	mergePRSrv MergePRSrv,
	reassignPRSrv ReassignPRSrv,
) Server {
	return Server{
		addTeamSrv: addTeamSrv,
		getTeamSrv: getTeamSrv,
		setIsActiveSrv: setIsActiveSrv,
		getUserReviewPRsSrv: getUserReviewPRsSrv,
		createPRSrv: createPRSrv,
		mergePRSrv: mergePRSrv,
		reassignPRSrv: reassignPRSrv,
	}
}
