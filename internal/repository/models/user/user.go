package user

import (
	"github.com/dexxhawk/pr-private/internal/domain"
)

type User struct {
	ID       string `db:"id"`
	Name     string `db:"name"`
	IsActive bool   `db:"is_active"`
	TeamName string `db:"team_name"`
}

func (User) Model(domain domain.User) User {
	return User{
		ID:       domain.ID,
		Name:     domain.Name,
		IsActive: domain.IsActive,
		TeamName: domain.TeamName,
	}
}

func (User) Models(domains ...domain.User) []User {
	models := make([]User, 0, len(domains))
	for _, d := range domains {
		models = append(models, User{}.Model(d))
	}

	return models
}

func (model User) Domain() domain.User {
	return domain.User{
		ID:       model.ID,
		Name:     model.Name,
		IsActive: model.IsActive,
		TeamName: model.TeamName,
	}
}

func (User) Domains(models []User) []domain.User {
	domains := make([]domain.User, 0, len(models))
	for _, m := range models {
		domains = append(domains, m.Domain())
	}
	return domains
}
