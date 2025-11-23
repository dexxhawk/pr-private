package set_isactive

type Service struct {
	userRepo UserRepo
}

func New(
	userRepo UserRepo,
) Service {
	return Service{
		userRepo: userRepo,
	}
}
