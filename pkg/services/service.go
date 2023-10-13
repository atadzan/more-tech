package services

import (
	"context"
	"github.com/atadzan/more-tech/models"
	"github.com/atadzan/more-tech/pkg/repository"
)

type IAtm interface {
	GetAll(ctx context.Context, page int) ([]models.Atm, error)
	GetById(ctx context.Context, id int) (models.AtmDetailedInfo, error)
}

type IOffice interface {
	GetAll(ctx context.Context, page int) ([]models.Office, error)
	GetById(ctx context.Context, id int) (models.OfficeDetailedInfo, error)
}

type Service struct {
	Atm    IAtm
	Office IOffice
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Atm:    NewAtmService(repo),
		Office: NewOfficeService(repo),
	}
}
