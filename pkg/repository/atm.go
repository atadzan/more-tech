package repository

import (
	"context"
	"github.com/atadzan/more-tech/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type AtmRepo struct {
	db *pgxpool.Pool
}

func NewAtmRepo(dbConn *pgxpool.Pool) *AtmRepo {
	return &AtmRepo{
		db: dbConn,
	}
}

func (r *AtmRepo) GetAll(ctx context.Context, page int) ([]models.Atm, error) {
	return nil, nil
}
func (r *AtmRepo) GetById(ctx context.Context, id int) (models.AtmDetailedInfo, error) {
	return models.AtmDetailedInfo{}, nil
}
