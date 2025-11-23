package team

import "github.com/dexxhawk/pr-private/internal/domain"

type Team struct {
	Name string `db:"name"`
}

func (Team) Model(domain domain.Team) Team {
	return Team{
		Name: domain.Name,
	}
}

func (model Team) Domain() domain.Team {
	return domain.Team{
		Name: model.Name,
	}
}
