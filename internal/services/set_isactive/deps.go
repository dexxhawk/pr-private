package set_isactive

import (
	"context"

	models_user "github.com/dexxhawk/pr-private/internal/repository/models/user"
)

type UserRepo interface {
	SetIsActive(ctx context.Context, userID string, isActive bool) (*models_user.User, error)
}
