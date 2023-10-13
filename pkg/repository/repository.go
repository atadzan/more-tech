package repository

import (
	"context"
	"github.com/atadzan/more-tech/models"
)

type IAtm interface {
	GetAll(ctx context.Context, page int) ([]models.Atm, error)
	GetById(ctx context.Context, id int) (models.AtmDetailedInfo, error)
}

type IOffice interface {
	GetAll(ctx context.Context, page int) ([]models.Office, error)
	GetById(ctx context.Context, id int) (models.OfficeDetailedInfo, error)
}

type Repository struct {
	Atm    IAtm
	Office IOffice
}

func NewRepository() *Repository {
	return &Repository{}
}
