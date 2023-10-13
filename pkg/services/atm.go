package services

import (
	"context"
	"github.com/atadzan/more-tech/models"
	"github.com/atadzan/more-tech/pkg/repository"
)

type Atm struct {
	repo repository.IAtm
}

func NewAtmService(repo repository.Repository) *Atm {
	return &Atm{
		repo: repo.Atm,
	}
}

func (a *Atm) GetAll(ctx context.Context, page int) ([]models.Atm, error) {
	return nil, nil
}
func (a *Atm) GetById(ctx context.Context, id int) (models.AtmDetailedInfo, error) {
	return models.AtmDetailedInfo{}, nil
}
