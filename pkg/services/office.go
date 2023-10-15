package services

import (
	"context"
	"github.com/atadzan/more-tech/models"
	"github.com/atadzan/more-tech/pkg/repository"
)

type Office struct {
	repo repository.IOfficeRepo
}

func NewOfficeService(repo *repository.Repository) *Office {
	return &Office{
		repo: repo.Office,
	}
}

func (s *Office) GetAll(ctx context.Context) ([]models.Office, error) {
	return s.repo.GetAll(ctx)
}
func (s *Office) GetById(ctx context.Context, id int) (models.OfficeDetailedInfo, error) {
	return s.repo.GetById(ctx, id)
}
