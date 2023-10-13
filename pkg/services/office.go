package services

import (
	"context"
	"github.com/atadzan/more-tech/models"
	"github.com/atadzan/more-tech/pkg/repository"
)

type Office struct {
	repo repository.IOffice
}

func NewOfficeService(repo repository.Repository) *Atm {
	return &Atm{
		repo: repo.Office,
	}
}

func (o *Office) GetAll(ctx context.Context, page int) ([]models.Office, error) {
	return nil, nil
}
func (o *Office) GetById(ctx context.Context, id int) (models.OfficeDetailedInfo, error) {
	return models.OfficeDetailedInfo{}, nil
}
