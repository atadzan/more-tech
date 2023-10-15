package services

import (
	"context"
	"github.com/atadzan/more-tech/models"
	"github.com/atadzan/more-tech/pkg/repository"
)

type Atm struct {
	repo repository.IAtmRepo
}

func NewAtmService(repo *repository.Repository) *Atm {
	return &Atm{
		repo: repo.Atm,
	}
}

func (s *Atm) GetAll(ctx context.Context) ([]models.Atm, error) {
	return s.repo.GetAll(ctx)
}
func (s *Atm) GetById(ctx context.Context, id int) (models.AtmDetailedInfo, error) {
	return models.AtmDetailedInfo{}, nil
}
