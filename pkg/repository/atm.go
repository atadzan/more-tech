package repository

import (
	"context"
	"fmt"
	"github.com/atadzan/more-tech/models"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

type AtmRepo struct {
	db *pgxpool.Pool
}

func NewAtmRepo(dbConn *pgxpool.Pool) *AtmRepo {
	return &AtmRepo{
		db: dbConn,
	}
}

func (r *AtmRepo) GetAll(ctx context.Context) (atms []models.Atm, err error) {
	query := fmt.Sprintf(`SELECT id, address, latitude, longitude, all_day FROM atms`)
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	for rows.Next() {
		var atm models.Atm
		if err = rows.Scan(&atm.Id, &atm.Address, &atm.Latitude, &atm.Longitude, &atm.AllDay); err != nil {
			log.Println(err)
			return nil, err
		}
		atms = append(atms, atm)
	}
	return
}
func (r *AtmRepo) GetById(ctx context.Context, id int) (atm models.AtmDetailedInfo, err error) {
	query := fmt.Sprintf(`SELECT id, address, latitude, longitude FROM atms WHERE id=$1`)
	if err = r.db.QueryRow(ctx, query, id).Scan(&atm.Id, &atm.Address, &atm.Latitude, &atm.Longitude); err != nil {
		log.Println(err)
		return models.AtmDetailedInfo{}, err
	}
	return
}
