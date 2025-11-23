package get_team

type Service struct {
	userRepo UserRepo
	teamRepo TeamRepo
}

func New(
	userRepo UserRepo,
	teamRepo TeamRepo,
) Service {
	return Service{
		userRepo: userRepo,
		teamRepo: teamRepo,
	}
}
