package user

import (
	"errors"
)

var ErrUserNotFound = errors.New("user not found")
var ErrTeamUsersNotFound = errors.New("users with such team not found")