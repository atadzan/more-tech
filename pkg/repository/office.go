package repository

import (
	"context"
	"fmt"
	"github.com/atadzan/more-tech/models"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

type OfficeRepo struct {
	db *pgxpool.Pool
}

func NewOfficeRepo(dbConn *pgxpool.Pool) *OfficeRepo {
	return &OfficeRepo{
		db: dbConn,
	}
}
func (r *OfficeRepo) GetAll(ctx context.Context) (offices []models.Office, err error) {
	query := fmt.Sprintf(`SELECT id, sale_point, type, sale_point_format, suo_availability, has_ramp, latitude, 
       longitude FROM offices`)
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	for rows.Next() {
		var office models.Office
		if err = rows.Scan(&office.Id, &office.SalePointName, &office.OfficeType, &office.SalePointFormat, &office.SuoAvailability,
			&office.HasRamp, &office.Latitude, &office.Longitude); err != nil {
			log.Println(err)
			return nil, err
		}
		offices = append(offices, office)
	}
	return
}
func (r *OfficeRepo) GetById(ctx context.Context, id int) (office models.OfficeDetailedInfo, err error) {
	query := fmt.Sprintf(`SELECT id, sale_point, type, sale_point_format, suo_availability, has_ramp, latitude, 
       longitude FROM offices  WHERE id = $1`)
	if err = r.db.QueryRow(ctx, query, id).Scan(&office.Id, &office.SalePointName, &office.OfficeType, &office.SalePointFormat, &office.SuoAvailability,
		&office.HasRamp, &office.Latitude, &office.Longitude); err != nil {
		log.Println(err)
		return models.OfficeDetailedInfo{}, err
	}

	return
}
