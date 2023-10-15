package repository

import (
	"context"
	"github.com/atadzan/more-tech/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type IAtmRepo interface {
	GetAll(ctx context.Context, page int) ([]models.Atm, error)
	GetById(ctx context.Context, id int) (models.AtmDetailedInfo, error)
}

type IOfficeRepo interface {
	GetAll(ctx context.Context, page int) ([]models.Office, error)
	GetById(ctx context.Context, id int) (models.OfficeDetailedInfo, error)
}

type Repository struct {
	Atm    IAtmRepo
	Office IOfficeRepo
}

func NewRepository(dbPoolConn *pgxpool.Pool) *Repository {
	return &Repository{
		Atm:    NewAtmRepo(dbPoolConn),
		Office: NewOfficeRepo(dbPoolConn),
	}
}
