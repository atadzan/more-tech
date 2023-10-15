package repository

import (
	"context"
	"github.com/atadzan/more-tech/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type OfficeRepo struct {
	db *pgxpool.Pool
}

func NewOfficeRepo(dbConn *pgxpool.Pool) *OfficeRepo {
	return &OfficeRepo{
		db: dbConn,
	}
}
func (r *OfficeRepo) GetAll(ctx context.Context, page int) ([]models.Office, error) {
	return nil, nil
}
func (r *OfficeRepo) GetById(ctx context.Context, id int) (models.OfficeDetailedInfo, error) {
	return models.OfficeDetailedInfo{}, nil
}
